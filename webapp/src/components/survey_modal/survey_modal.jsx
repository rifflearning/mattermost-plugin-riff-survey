import React from 'react';
import PropTypes from 'prop-types';

import {Alert, Button, ButtonGroup, Clearfix, Modal} from 'react-bootstrap';

import QuestionTypeOpen from '../question_type_open';
import QuestionTypeLikertScale from '../question_type_likert_scale';

import constants from '../../constants';
import loadingGif from '../../../assets/load.gif';

import './styles.css';

export default class SurveyModal extends React.PureComponent {
    static propTypes = {
        theme: PropTypes.object.isRequired,
        surveyPostID: PropTypes.string.isRequired,
        surveyPostProps: PropTypes.object.isRequired,
        visible: PropTypes.bool.isRequired,
        close: PropTypes.func.isRequired,
        getSurvey: PropTypes.func.isRequired,
        submitSurveyResponses: PropTypes.func.isRequired,
    };

    constructor(props) {
        super(props);
        this.state = {
            survey: {
                id: '',
                version: '',
                title: '',
                description: '',
                questions: [],
            },
            responses: {
            },
            loading: false,
            loadingSubmit: true,
            serverError: false,
        };
    }

    componentDidUpdate(prevProps) {
        if (this.props.visible && !prevProps.visible) {
            this.getSurvey();
        }
    }

    getSurvey = async () => {
        const {surveyPostProps, getSurvey} = this.props;

        // TODO: Get survey using meetingID instead
        const {data} = await getSurvey(surveyPostProps.survey_id, surveyPostProps.survey_version);
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
    };

    handleClose = () => {
        this.props.close();
    };

    handleSubmit = () => {
        const {survey, responses} = this.state;
        const {surveyPostProps, surveyPostID} = this.props;
        const meetingID = surveyPostProps.meeting_id;
        this.props.submitSurveyResponses(surveyPostID, meetingID, survey.id, survey.version, responses);
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

    renderLoading = () => {
        return (
            <div className='survey-loading'>
                <img src={loadingGif}/>
            </div>
        );
    };

    renderSurvey = () => {
        const {survey, loadingSubmit, serverError} = this.state;

        const questions = this.renderQuestions();
        return (
            <div>
                <p className='survey-banner-text'>
                    {survey.description}
                </p>
                {questions}
                <Clearfix>
                    <ButtonGroup className='float-right'>
                        <Button
                            type='button'
                            bsStyle='info'
                            onClick={this.handleClose}
                            disabled={loadingSubmit}
                        >
                            {'Cancel'}
                        </Button>
                        <Button
                            type='submit'
                            bsStyle='primary'
                            className='submit-survey-btn'
                            onClick={this.handleSubmit}
                            disabled={loadingSubmit}
                        >
                            {loadingSubmit && (
                                <span
                                    className='fa fa-spinner fa-fw fa-pulse spinner'
                                    title={'Loading Icon'}
                                />
                            )}
                            {'Submit'}
                        </Button>
                    </ButtonGroup>
                </Clearfix>
                {serverError && (
                    <Alert
                        bsStyle='warning'
                        className='survey-server-error-alert'
                    >
                        <i
                            className='fa fa-warning'
                            title='Server Error'
                        />
                        {' There was some error while submitting your response. Please try again later. If the problem persists, contact your System Administrator.'}
                    </Alert>
                )}
            </div>
        );
    };

    renderGetSurveyError = () => {
        return (
            <div>
                <i
                    className='fa fa-warning'
                    title='Server Error'
                />
                {' There was some error while fetching survey. Please try again later. If the problem persists, contact your System Administrator.'}
            </div>
        );
    };

    render() {
        const {survey, loading, loadingSubmit, serverError} = this.state;

        let content;
        // TODO: Set the value of content

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
                    {content}
                </Modal.Body>
            </Modal>
        );
    }
}
