import React from 'react';
import PropTypes from 'prop-types';

import {Button, ButtonGroup, Modal} from 'react-bootstrap';

import QuestionTypeOpen from '../question_type_open';
import QuestionTypeLikertScale from '../question_type_likert_scale';

import constants from '../../constants';

import './styles.css';

export default class SurveyModal extends React.PureComponent {
    static propTypes = {
        theme: PropTypes.object.isRequired,
        currentPostProps: PropTypes.object.isRequired,
        visible: PropTypes.bool.isRequired,
        close: PropTypes.func.isRequired,
        getSurvey: PropTypes.func.isRequired,
        submitSurveyResponses: PropTypes.func.isRequired,
    }

    constructor(props) {
        super(props);
        this.state = {
            survey: {
                title: '',
                description: '',
                questions: [],
            },
            responses: {
            },
        };
    }

    componentDidUpdate(prevProps) {
        if (this.props.visible && !prevProps.visible) {
            this.getSurvey();
        }
    }

    getSurvey = async () => {
        const {currentPostProps, getSurvey} = this.props;

        // TODO: Get survey using meetingID instead
        const {data} = await getSurvey(currentPostProps.survey_id, currentPostProps.survey_version);
        if (data) {
            const survey = data;
            const responses = survey.questions.reduce((obj, question) => {
                obj[question.id] = '';
                return obj;
            }, {});

            this.setState({
                survey,
                responses,
            });
        }
    }

    handleClose = () => {
        this.props.close();
    };

    handleSubmit = () => {
        const {responses} = this.state;
        const {currentPostProps} = this.props;
        const meetingID = currentPostProps.meeting_id;
        const surveyID = currentPostProps.survey_id;
        const surveyVersion = currentPostProps.survey_version;
        this.props.submitSurveyResponses(meetingID, surveyID, surveyVersion, responses);
        this.handleClose();
    };

    handleUpdateQuestionResponse = (questionID, response) => {
        this.setState((prevState) => {
            const responses = {...prevState.responses};
            responses[questionID] = response;
            return {
                responses,
            };
        });
    };

    renderQuestions = () => {
        const {theme} = this.props;
        const questionsList = this.state.survey.questions;

        return questionsList.map((question, idx) => {
            const baseProps = {
                index: idx + 1,
                id: question.id,
                key: question.id,
                text: question.text,
                theme,
                handleChange: this.handleUpdateQuestionResponse,
            };
            switch (question.type) {
            case constants.QUESTION_TYPES.OPEN:
                return (
                    <QuestionTypeOpen
                        {...baseProps}
                    />
                );

            case constants.QUESTION_TYPES.FIVE_POINT_LIKERT_SCALE:
                return (
                    <QuestionTypeLikertScale
                        {...baseProps}
                        responses={constants.FIVE_POINT_LIKERT_SCALE_RESPONSES}
                    />
                );

            default:
                return null;
            }
        });
    };

    render() {
        const {survey} = this.state;
        const questions = this.renderQuestions();
        return (
            <Modal
                show={this.props.visible}
                onHide={this.handleClose}
                backdrop={'static'}
                centered={true}
            >
                <Modal.Header
                    closeButton={true}
                    closeLabel={'Close'}
                >
                    <Modal.Title>
                        {survey.title}
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <p className='survey-banner-text'>
                        {survey.description}
                    </p>
                    {questions}
                    <ButtonGroup className='float-right'>
                        <Button
                            type='button'
                            bsStyle='secondary'
                            onClick={this.handleClose}
                        >
                            {'Cancel'}
                        </Button>
                        <Button
                            type='submit'
                            bsStyle='primary'
                            className='submit-survey-btn'
                            onClick={this.handleSubmit}
                        >
                            {'Submit'}
                        </Button>
                    </ButtonGroup>
                </Modal.Body>
            </Modal>
        );
    }
}
