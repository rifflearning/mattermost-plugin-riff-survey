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
            hoveredValue: e.target.dataset.value,
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
                <label
                    key={index + response.value}
                    className='form-check'
                    style={optionStyle}
                    htmlFor={`${index}${idx}`}
                    data-value={response.value}
                    onMouseOver={this.handleMouseEnter}
                    onMouseOut={this.handleMouseLeave}
                >
                    <input
                        type='radio'
                        className='form-check-input'
                        value={response.value}
                        name={index}
                        id={`${index}${idx}`}
                        style={style.optionRadio}
                        onClick={this.handleChange}
                    />
                    <span
                        className='form-check-label'
                        style={style.optionLabel}
                    >
                        {response.text}
                    </span>
                </label>
            );
        });

        return (
            <div className='form-group'>
                <p>{`${index}. ${text}`}</p>
                <div style={style.options}>{radios}</div>
            </div>
        );
    }
}

const style = {
    options: {
        marginTop: '20px',
        marginBottom: '32px',
        width: '100%',
        height: '64px',
        display: 'flex',
        flexDirection: 'row',
        borderWidth: '1px',
        borderStyle: 'solid',
        borderImage: 'initial',
        borderRadius: '3px',
        overflow: 'hidden',
    },
    option: {
        height: '100%',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        cursor: 'pointer',
        userSelect: 'none',
        flex: '1 1 0%',
    },
    optionRadio: {
        display: 'none',
    },
    optionLabel: {
        display: 'inline-block',
        cursor: 'pointer',
        textAlign: 'center',
    },
};
