import Constants from '../constants';

export const active = (state = false, action) => {
    switch (action.type) {
    case Constants.ACTION_TYPES.PLUGIN_ENABLED:
        return true;
    case Constants.ACTION_TYPES.PLUGIN_DISABLED:
        return false;
    default:
        return state;
    }
};
