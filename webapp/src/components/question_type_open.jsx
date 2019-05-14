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

        return (
            <div className='form-group clearfix'>
                <p>{`${index}. ${text}`}</p>
                <textarea
                    maxLength={Constants.OPEN_QUESTION_MAX_LENGTH}
                    onChange={this.handleChange}
                    className='form-control'
                    rows={5}
                    style={style.textarea}
                    id={index}
                />
                <span style={style.remaining}>{`${this.state.remaining} character(s) left`}</span>
            </div>
        );
    }
}

const style = {
    remaining: {
        float: 'right',
    },
    textarea: {
        resize: 'vertical',
    },
};
