import Client from '../client';
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

export const setCurrentPostProps = (data) => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.SET_CURRENT_POST_PROPS,
        data,
    });
};

export function getSurvey(surveyID, surveyVersion) {
    return async () => {
        let data;
        try {
            data = await Client.getSurvey(surveyID, surveyVersion);
        } catch (error) {
            return {data: null, error};
        }

        return {data, error: null};
    };
}
