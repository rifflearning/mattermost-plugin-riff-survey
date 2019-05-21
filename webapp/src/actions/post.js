import Constants from '../constants';

export const setCurrentPostProps = (data) => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.SET_CURRENT_POST_PROPS,
        data,
    });
};
