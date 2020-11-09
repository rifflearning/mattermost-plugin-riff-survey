import {combineReducers} from 'redux';

import {active} from './active';
import {survey} from './survey';

export default combineReducers({
    active,
    survey,
});
