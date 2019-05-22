import Constants from '../constants';

export const setCurrentPostID = (data) => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.SET_CURRENT_POST_ID,
        data,
    });
};

export const setCurrentPostProps = (data) => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.SET_CURRENT_POST_PROPS,
        data,
    });
};
