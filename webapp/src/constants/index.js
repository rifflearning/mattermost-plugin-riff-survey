import {PLUGIN_NAME} from './manifest';
import {ACTION_TYPES} from './action_types';

const OPEN_QUESTION_INITIAL_ROWS = 5;
const OPEN_QUESTION_MAX_LENGTH = 2000;

const QUESTION_TYPES = {
    OPEN: 'open',
    FIVE_POINT_LIKERT_SCALE: 'five-point-likert-scale',
};

//
// Export the constants
//
export default {
    ACTION_TYPES,
    OPEN_QUESTION_INITIAL_ROWS,
    OPEN_QUESTION_MAX_LENGTH,
    PLUGIN_NAME,
    QUESTION_TYPES,
};
