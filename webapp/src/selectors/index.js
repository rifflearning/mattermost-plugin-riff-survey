import Constants from '../constants';

const getPluginState = (state) => state[`plugins-${Constants.PLUGIN_NAME}`] || {};

const isSurveyModalVisible = (state) => getPluginState(state).surveyModalVisible || false;

const currentPostProps = (state) => getPluginState(state).currentPostProps || {};

const dashboardPath = (state) => getPluginState(state).dashboardPath || '';

export default {
    currentPostProps,
    dashboardPath,
    isSurveyModalVisible,
};
