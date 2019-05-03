import request from 'superagent';

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

    exampleRequest = async () => {
        return this.doGet(`${this.url}/example`);
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
