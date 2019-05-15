import React from 'react';
import PropTypes from 'prop-types';

import {Button, ButtonGroup, Modal} from 'react-bootstrap';

import QuestionTypeOpen from '../question_type_open';
import QuestionTypeLikertScale from '../question_type_likert_scale';

import './styles.css';

const questionsList = [
    {
        type: '5-point-likert-scale',
        text: 'I felt comfortable conversing using this medium.',
    },
    {
        type: '5-point-likert-scale',
        text: 'I felt comfortable participating in group discussions.',
    },
    {
        type: '5-point-likert-scale',
        text: 'I felt comfortable interacting with other group members.',
    },
    {
        type: '5-point-likert-scale',
        text: 'I was able to speak my mind in my group discussion.',
    },
    {
        type: 'open',
        text: 'Please add any other comments about the Riff Edu meeting experience.',
    },
];

export default class SurveyModal extends React.PureComponent {
    static propTypes = {
        theme: PropTypes.object.isRequired,
        visible: PropTypes.bool,
        close: PropTypes.func,
    }

    constructor(props) {
        super(props);
        this.state = {
        };
    }

    handleClose = () => {
        this.props.close();
    };

    handleSubmit = () => {
        // TODO: API calls
        this.handleClose();
    };

    renderQuestions = () => {
        const {theme} = this.props;

        return questionsList.map((question, idx) => {
            switch (question.type) {
            case 'open':
                return (
                    <QuestionTypeOpen
                        index={idx + 1}
                        key={question.text}
                        text={question.text}
                        theme={theme}
                    />
                );

            case '5-point-likert-scale':
                return (
                    <QuestionTypeLikertScale
                        index={idx + 1}
                        key={question.text}
                        text={question.text}
                        theme={theme}
                    />
                );

            default:
                return null;
            }
        });
    };
    render() {
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
                        {'Riff Meeting Survey'}
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <p className='survey-banner-text'>
                        {'Please tell us about your Riff meeting experience. We will ask you to take this short survey after each meeting, to see how your experience changes over time.'}
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
