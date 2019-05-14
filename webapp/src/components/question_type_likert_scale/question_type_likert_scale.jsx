import React from 'react';
import PropTypes from 'prop-types';

import {FormGroup, ControlLabel, Radio} from 'react-bootstrap';
import {changeOpacity} from 'mattermost-redux/utils/theme_utils';

import './styles.css';

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
            const optionStyle = {};
            if (selectedValue === response.value) {
                optionStyle.backgroundColor = this.props.theme.sidebarTextActiveBorder;
                optionStyle.color = this.props.theme.sidebarTextActiveColor;
            } else if (hoveredValue === response.value) {
                optionStyle.backgroundColor = changeOpacity(this.props.theme.sidebarTextHoverBg, 0.1);
            }

            return (
                <ControlLabel
                    key={index + response.value}
                    className='likert-option'
                    style={optionStyle}
                    htmlFor={`${index}${idx}`}
                    data-value={response.value}
                    onMouseOver={this.handleMouseEnter}
                    onMouseOut={this.handleMouseLeave}
                >
                    <Radio
                        className='display-none'
                        value={response.value}
                        name={index}
                        id={`${index}${idx}`}
                        onClick={this.handleChange}
                    />
                    <span className='likert-option-label'>{response.text}</span>
                </ControlLabel>
            );
        });

        return (
            <FormGroup>
                <p>{`${index}. ${text}`}</p>
                <div className='likert-options'>{radios}</div>
            </FormGroup>
        );
    }
}
