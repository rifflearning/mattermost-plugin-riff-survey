import Constants from '../constants';

const getPluginState = (state) => state[`plugins-${Constants.PLUGIN_NAME}`] || {};

const isSurveyModalVisible = (state) => getPluginState(state).surveyModalVisible || false;

export default {
    isSurveyModalVisible,
};
