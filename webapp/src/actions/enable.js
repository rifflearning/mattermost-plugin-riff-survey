import Constants from '../constants';

export const pluginEnabled = () => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.PLUGIN_ENABLED,
    });
};

export const pluginDisabled = () => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.PLUGIN_DISABLED,
    });
};
