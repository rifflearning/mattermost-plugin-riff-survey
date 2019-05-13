import Constants from '../constants';

export const openSurveyModal = () => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.OPEN_SURVEY_MODAL,
    });
};

export const closeSurveyModal = () => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.CLOSE_SURVEY_MODAL,
    });
};
