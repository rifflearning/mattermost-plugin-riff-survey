import Constants from '../constants';

export const surveyModalVisible = (state = false, action) => {
    switch (action.type) {
    case Constants.ACTION_TYPES.OPEN_SURVEY_MODAL:
        return true;
    case Constants.ACTION_TYPES.CLOSE_SURVEY_MODAL:
        return false;
    default:
        return state;
    }
};
