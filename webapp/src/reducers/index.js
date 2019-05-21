import {combineReducers} from 'redux';

import {currentPostProps, surveyModalVisible} from './survey';

export default combineReducers({
    currentPostProps,
    surveyModalVisible,
});
