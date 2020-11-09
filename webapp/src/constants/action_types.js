import {PLUGIN_NAME} from './manifest';

export const ACTION_TYPES = {
    PLUGIN_ENABLED: `${PLUGIN_NAME}_plugin_enabled`,
    PLUGIN_DISABLED: `${PLUGIN_NAME}_plugin_disabled`,
    OPEN_SURVEY_MODAL: `${PLUGIN_NAME}_open_survey_modal`,
    CLOSE_SURVEY_MODAL: `${PLUGIN_NAME}_close_survey_modal`,
    SURVEY_SUBMIT_SUCCESS: `${PLUGIN_NAME}_survey_submit_success`,

    // From mattermost-plugin-riff-core.
    OPEN_RIFF_DASHBOARD: 'ai.riffanalytics.core_OPEN_RIFF_METRICS_MODAL',
};
