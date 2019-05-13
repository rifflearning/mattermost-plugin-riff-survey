import React from 'react';
import PropTypes from 'prop-types';

import {Button} from 'react-bootstrap';

import {formatText, messageHtmlToComponent} from 'post-utils';

export default class PostTypeSurvey extends React.PureComponent {
    static propTypes = {
        post: PropTypes.object.isRequired,
        open: PropTypes.func,
        // eslint-disable-next-line lines-around-comment
        // theme: PropTypes.object.isRequired,
    }

    constructor(props) {
        super(props);
        this.state = {
        };
    }

    openModal = () => {
        this.props.open();
    };

    render() {
        const post = {...this.props.post};
        const message = post.message || '';
        return (
            <React.Fragment>
                {messageHtmlToComponent(formatText(message))}
                <Button onClick={this.openModal}>{'Click here'}</Button>
            </React.Fragment>
        );
    }
}
