import React from 'react';
import PropTypes from 'prop-types';

import {changeOpacity} from 'mattermost-redux/utils/theme_utils';

export default class QuestionTypeLikertScale extends React.PureComponent {
    static propTypes = {
        text: PropTypes.string.isRequired,
        theme: PropTypes.object.isRequired,
        // eslint-disable-next-line react/no-unused-prop-types
        points: PropTypes.number,
        index: PropTypes.number,
        responses: PropTypes.array,
        handleChange: PropTypes.func,
    }

    static defaultProps = {
        points: 5,
        responses: [
            {value: '1', text: 'Strongly Agree'},
            {value: '2', text: 'Agree'},
            {value: '3', text: 'Neutral'},
            {value: '4', text: 'Disagree'},
            {value: '5', text: 'Strongly Disagree'},
        ],
        handleChange: (val) => {
            console.log(val); // eslint-disable-line no-console
        },
    };

    constructor(props) {
        super(props);
        this.state = {
            selectedValue: '',
            hoveredValue: '',
        };
    }

    handleChange = (evt) => {
        const selectedValue = evt.target.value;
        this.props.handleChange(selectedValue);
        this.setState({
            selectedValue,
        });
    };

    handleMouseEnter = (e) => {
        this.setState({
            hoveredValue: e.target.value,
        });
    }

    handleMouseLeave = () => {
        this.setState({
            hoveredValue: '',
        });
    }

    render() {
        const {index, responses, text} = this.props;
        const {hoveredValue, selectedValue} = this.state;

        const radios = responses.map((response, idx) => {
            const optionStyle = {...style.option};
            if (selectedValue === response.value) {
                optionStyle.backgroundColor = this.props.theme.sidebarTextActiveBorder;
                optionStyle.color = this.props.theme.sidebarTextActiveColor;
            } else if (hoveredValue === response.value) {
                optionStyle.backgroundColor = changeOpacity(this.props.theme.sidebarTextHoverBg, 0.1);
            }

            return (
                <div
                    key={index + response.value}
                    className='form-check'
                    style={optionStyle}
                >
                    <input
                        type='radio'
                        className='form-check-input'
                        value={response.value}
                        name={text}
                        id={text + idx}
                        onMouseEnter={this.handleMouseEnter}
                        onMouseLeave={this.handleMouseLeave}
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
            <div style={style.question}>
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
    questionText: {
        display: 'block',
        width: '100%',
        padding: '0',
        color: '#333',
    },
    options: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-between',
        paddingTop: '1em',
        cursor: 'pointer',
    },
    option: {
        display: 'flex',
        flexDirection: 'column',
        flex: '1 1 auto',
        alignItems: 'center',
        margin: '0 10px',
    },
    optionText: {
        display: 'inline-block',
        paddingTop: '0.4em',
    },
};
