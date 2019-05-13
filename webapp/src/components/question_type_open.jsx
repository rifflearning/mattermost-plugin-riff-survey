import React from 'react';
import PropTypes from 'prop-types';

import Constants from '../constants';

export default class QuestionTypeOpen extends React.PureComponent {
    static propTypes = {
        text: PropTypes.string.isRequired,
        index: PropTypes.number,
    }

    constructor(props) {
        super(props);
        this.state = {
            remaining: Constants.OPEN_QUESTION_MAX_LENGTH,
        };
    }

    handleChange = (e) => {
        this.setState({
            remaining: Constants.OPEN_QUESTION_MAX_LENGTH - e.target.value.length,
        });
    };

    render() {
        const {index, text} = this.props;

        let questionStyles = style.question;
        if (index % 2 === 1) {
            questionStyles = {...questionStyles, ...style.primaryQuestion};
        }

        return (
            <fieldset style={questionStyles}>
                <span style={style.questionText}>{`${index}. ${text}`}</span>
                <textarea
                    maxLength={Constants.OPEN_QUESTION_MAX_LENGTH}
                    onChange={this.handleChange}
                    className='form-control'
                    rows={5}
                />
                <span style={style.remaining}>{`${this.state.remaining} character(s) left`}</span>
            </fieldset>
        );
    }
}

const style = {
    question: {
        margin: '1em 0',
    },
    primaryQuestion: {
        backgroundColor: '#7e7e7e',
    },
    questionText: {
        display: 'block',
        width: '100%',
        padding: '0',
        fontSize: '1.2em',
        color: '#333',
    },
    remaining: {
        float: 'right',
    },
};
