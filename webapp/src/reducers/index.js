import {combineReducers} from 'redux';

import {currentPostProps, dashboardPath, surveyModalVisible} from './survey';

export default combineReducers({
    currentPostProps,
    dashboardPath,
    surveyModalVisible,
});
