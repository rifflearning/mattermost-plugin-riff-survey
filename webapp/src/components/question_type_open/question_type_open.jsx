import React from 'react';
import PropTypes from 'prop-types';
import {FormGroup, FormControl, HelpBlock} from 'react-bootstrap';

import Constants from '../../constants';
import './styles.css';

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

        return (
            <FormGroup className='clearfix'>
                <p>{`${index}. ${text}`}</p>
                <FormControl
                    componentClass='textarea'
                    maxLength={Constants.OPEN_QUESTION_MAX_LENGTH}
                    onChange={this.handleChange}
                    rows={Constants.OPEN_QUESTION_INITIAL_ROWS}
                    className='open-question-textarea'
                    id={index}
                />
                <HelpBlock className='float-right'>{`${this.state.remaining} character(s) left`}</HelpBlock>
            </FormGroup>
        );
    }
}