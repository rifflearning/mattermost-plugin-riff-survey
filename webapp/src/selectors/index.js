import Constants from '../constants';

const getPluginState = (state) => state[`plugins-${Constants.PLUGIN_NAME}`] || {};

const isSurveyModalVisible = (state) => getPluginState(state).surveyModalVisible || false;

const currentPostID = (state) => getPluginState(state).currentPostID || '';

const currentPostProps = (state) => getPluginState(state).currentPostProps || {};

const dashboardPath = (state) => getPluginState(state).dashboardPath || '';

export default {
    currentPostID,
    currentPostProps,
    dashboardPath,
    isSurveyModalVisible,
};
