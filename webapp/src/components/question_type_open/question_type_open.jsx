import React from 'react';
import PropTypes from 'prop-types';
import {FormGroup, FormControl} from 'react-bootstrap';

import Constants from '../../constants';
import './styles.css';

export default class QuestionTypeOpen extends React.PureComponent {
    static propTypes = {
        id: PropTypes.string.isRequired,
        text: PropTypes.string.isRequired,
        index: PropTypes.number.isRequired,
        handleChange: PropTypes.func.isRequired,
    };

    constructor(props) {
        super(props);
        this.state = {
            remaining: Constants.OPEN_QUESTION_MAX_LENGTH,
        };
    }

    handleChange = (e) => {
        const value = e.target.value;
        this.props.handleChange(this.props.id, value);
        this.setState({
            remaining: Constants.OPEN_QUESTION_MAX_LENGTH - value.length,
        });
    };

    render() {
        const {id: questionID, index, text} = this.props;

        return (
            <FormGroup className='clearfix'>
                <p id={questionID}>{`${index}. ${text}`}</p>
                <FormControl
                    aria-labelledby={questionID}
                    componentClass='textarea'
                    className='open-question-textarea'
                    onChange={this.handleChange}
                    rows={Constants.OPEN_QUESTION_INITIAL_ROWS}
                    maxLength={Constants.OPEN_QUESTION_MAX_LENGTH}
                />
                <p className='float-right'>{`${this.state.remaining} character(s) left`}</p>
            </FormGroup>
        );
    }
}
