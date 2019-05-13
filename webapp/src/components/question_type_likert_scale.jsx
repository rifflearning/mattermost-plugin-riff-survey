import React from 'react';
import PropTypes from 'prop-types';

export default class QuestionTypeLikertScale extends React.PureComponent {
    static propTypes = {
        text: PropTypes.string.isRequired,
        // eslint-disable-next-line react/no-unused-prop-types
        points: PropTypes.number,
        index: PropTypes.number,
        responses: PropTypes.array,
        handleChange: PropTypes.func,
    }

    static defaultProps = {
        points: 5,
        responses: [
            {value: 1, text: 'Strongly Agree'},
            {value: 2, text: 'Agree'},
            {value: 3, text: 'Neutral'},
            {value: 4, text: 'Disagree'},
            {value: 5, text: 'Strongly Disagree'},
        ],
        handleChange: (val) => {
            console.log(val); // eslint-disable-line no-console
        },
    };

    constructor(props) {
        super(props);
        this.state = {
        };
    }

    handleChange = (evt) => {
        this.props.handleChange(evt.target.value);
    };

    render() {
        const {index, responses, text} = this.props;

        let questionStyles = style.question;
        if (index % 2 === 1) {
            questionStyles = {...questionStyles, ...style.primaryQuestion};
        }

        const radios = responses.map((response, idx) => {
            return (
                <div
                    key={index + response.value}
                    className='form-check'
                    style={style.option}
                >
                    <input
                        type='radio'
                        className='form-check-input'
                        value={response.value}
                        name={text}
                        id={text + idx}
                        onClick={this.handleChange}
                    />
                    <label
                        className='form-check-label'
                        style={style.optionText}
                    >
                        {response.text}
                    </label>
                </div>
            );
        });

        return (
            <div style={questionStyles}>
                <span style={style.questionText}>{`${index}. ${text}`}</span>
                <div style={style.options}>{radios}</div>
            </div>
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
    options: {
        display: 'flex',
        flexDirection: 'row',
        paddingTop: '1em',
    },
    option: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        margin: '0 10px',
    },
    optionText: {
        display: 'inline-block',
        paddingTop: '0.4em',
    },
};
