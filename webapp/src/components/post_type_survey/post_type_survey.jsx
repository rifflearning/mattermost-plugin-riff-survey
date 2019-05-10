import React from 'react';
import PropTypes from 'prop-types';

import {formatText, messageHtmlToComponent} from 'post-utils';

export default class PostTypeSurvey extends React.PureComponent {
    static propTypes = {
        post: PropTypes.object.isRequired,
        theme: PropTypes.object.isRequired,
    }

    constructor(props) {
        super(props);
        this.state = {
        };
    }

    render() {
        const post = {...this.props.post};
        const message = post.message || '';
        return messageHtmlToComponent(formatText(message));
    }
}
