import React from 'react';
import PropTypes from 'prop-types';

import {formatText, messageHtmlToComponent} from 'post-utils';

import './styles.css';

export default class PostTypeSurvey extends React.PureComponent {
    static propTypes = {
        post: PropTypes.object.isRequired,
        currentUser: PropTypes.object.isRequired,
        open: PropTypes.func,
    }

    constructor(props) {
        super(props);
        this.state = {
        };
    }

    openModal = () => {
        this.props.open();
    };

    goToDashboard = () => {
        // TODO: open dashboard page
    };

    renderSubmitted = () => {
        const message1 = `Thanks for your feedback @${this.props.currentUser.username}! Have you checked out`;
        const dashboardLink = (
            <p><a onClick={this.goToDashboard}>{'your Riff Stats'}</a></p>
        );
        const message2 = ' in the Dashboard?';
        return (
            <div className='same-line'>
                {messageHtmlToComponent(formatText(message1, {atMentions: true}))}
                {dashboardLink}
                {messageHtmlToComponent(formatText(message2))}
            </div>
        );
    };

    renderNotSubmitted = () => {
        const message1 = `Hi @${this.props.currentUser.username} - Please`;
        const modalLink = (
            <p><a onClick={this.openModal}>{'tell us about the meeting'}</a></p>
        );

        // TODO: Get other meeting participants
        const message2 = ' you just had? It only takes 30 seconds, and helps understand your experience over time.';

        return (
            <div className='same-line'>
                {messageHtmlToComponent(formatText(message1, {atMentions: true}))}
                {modalLink}
                {messageHtmlToComponent(formatText(message2, {atMentions: true}))}
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
