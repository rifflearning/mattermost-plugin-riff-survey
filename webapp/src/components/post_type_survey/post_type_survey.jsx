import React from 'react';
import PropTypes from 'prop-types';

import {formatText, messageHtmlToComponent} from 'post-utils';

import {Button} from 'react-bootstrap';

import './styles.css';

export default class PostTypeSurvey extends React.PureComponent {
    static propTypes = {
        post: PropTypes.object.isRequired,
        currentUser: PropTypes.object.isRequired,
        setCurrentPostProps: PropTypes.func.isRequired,
        open: PropTypes.func.isRequired,
    }

    constructor(props) {
        super(props);
        this.state = {
        };
    }

    openModal = () => {
        this.props.setCurrentPostProps(this.props.post.props);
        this.props.open();
    };

    goToDashboard = () => {
        // TODO: open dashboard page
    };

    renderSubmitted = () => {
        const message = `Thanks for your feedback @${this.props.currentUser.username}! Have you checked out your Riff Stats in the Dashboard?`;
        return (
            <div>
                {messageHtmlToComponent(formatText(message, {atMentions: true}))}
                <Button
                    bsStyle='primary'
                    className='survey-action-button'
                    onClick={this.goToDashboard}
                >
                    {'Click Here'}
                </Button>
            </div>
        );
    };

    renderNotSubmitted = () => {
        const message = `Hi @${this.props.currentUser.username} - Please tell us about the meeting you just had? It only takes 30 seconds, and helps understand your experience over time.`;

        return (
            <div>
                {messageHtmlToComponent(formatText(message, {atMentions: true}))}
                <Button
                    bsStyle='primary'
                    className='survey-action-button'
                    onClick={this.openModal}
                >
                    {'Click Here'}
                </Button>
            </div>
        );
    };

    render() {
        const {post} = this.props;
        const postProps = post.props;
        if (postProps.submitted) {
            return this.renderSubmitted();
        }

        return this.renderNotSubmitted();
    }
}
