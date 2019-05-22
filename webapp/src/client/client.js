import request from 'superagent';
import {buildQueryString} from 'mattermost-redux/utils/helpers';

import Constants from '../constants';

/**
 *  Add web utilities for interacting with servers here
 */

export default class Client {
    constructor() {
        const url = new URL(window.location.href);
        this.baseUrl = `${url.protocol}//${url.host}`;
        this.pluginUrl = `${this.baseUrl}/plugins/${Constants.PLUGIN_NAME}`;
        this.apiUrl = `${this.baseUrl}/api/v4`;
    }

    getSurvey = async (surveyID, surveyVersion) => {
        const queryParams = {
            survey_id: surveyID,
            survey_version: surveyVersion,
        };
        const url = `${this.pluginUrl}/survey${buildQueryString(queryParams)}`;
        return this.doGet(url);
    };

    getDashboardURL = (teamName, path) => {
        let url = this.baseUrl;
        if (teamName && path) {
            url = `${this.baseUrl}/${teamName}${path}`;
        }
        return url;
    };

    getDashboardPath = () => {
        return this.doGet(`${this.pluginUrl}/dashboard`);
    };

    submitSurveyResponses = (meetingID, surveyID, surveyVersion, responses) => {
        const url = `${this.pluginUrl}/submit`;
        const body = {
            meeting_id: meetingID,
            survey_id: surveyID,
            survey_version: surveyVersion,
            responses,
        };
        return this.doPost(url, body);
    };

    doGet = async (url, headers = {}) => {
        headers['X-Requested-With'] = 'XMLHttpRequest';

        try {
            const response = await request.
                get(url).
                set(headers).
                type('application/json').
                accept('application/json');

            return response.body;
        } catch (err) {
            throw err;
        }
    };

    doPost = async (url, body, headers = {}) => {
        headers['X-Requested-With'] = 'XMLHttpRequest';

        try {
            const response = await request.
                post(url).
                send(body).
                set(headers).
                type('application/json').
                accept('application/json');

            return response.body;
        } catch (err) {
            throw err;
        }
    };

    doDelete = async (url, body, headers = {}) => {
        headers['X-Requested-With'] = 'XMLHttpRequest';

        try {
            const response = await request.
                delete(url).
                send(body).
                set(headers).
                type('application/json').
                accept('application/json');

            return response.body;
        } catch (err) {
            throw err;
        }
    };

    doPut = async (url, body, headers = {}) => {
        headers['X-Requested-With'] = 'XMLHttpRequest';

        try {
            const response = await request.
                put(url).
                send(body).
                set(headers).
                type('application/json').
                accept('application/json');

            return response.body;
        } catch (err) {
            throw err;
        }
    }
}
