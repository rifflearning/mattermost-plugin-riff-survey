import React from 'react';
import PropTypes from 'prop-types';

import {Button, Modal} from 'react-bootstrap';

import QuestionTypeOpen from '../question_type_open';
import QuestionTypeLikertScale from '../question_type_likert_scale';

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
        visible: PropTypes.bool,
        close: PropTypes.func,
        // eslint-disable-next-line lines-around-comment
        // theme: PropTypes.object.isRequired,
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
        return questionsList.map((question, idx) => {
            switch (question.type) {
            case 'open':
                return (
                    <QuestionTypeOpen
                        index={idx + 1}
                        key={question.text}
                        text={question.text}
                    />
                );

            case '5-point-likert-scale':
                return (
                    <QuestionTypeLikertScale
                        index={idx + 1}
                        key={question.text}
                        text={question.text}
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
                    <p>{'Please tell us about your Riff meeting experience. We will ask you to take this short survey after each meeting, to see how your experience changes over time.'}</p>
                    {questions}
                    <div
                        role='group'
                        className='btn-group'
                        style={{float: 'right'}}
                    >
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
                            onClick={this.handleSubmit}
                        >
                            {'Submit'}
                        </Button>
                    </div>
                </Modal.Body>
            </Modal>
        );
    }
}
