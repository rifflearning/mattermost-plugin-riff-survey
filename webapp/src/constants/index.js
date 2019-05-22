import {PLUGIN_NAME} from './manifest';
import {ACTION_TYPES} from './action_types';

const OPEN_QUESTION_INITIAL_ROWS = 5;
const OPEN_QUESTION_MAX_LENGTH = 2000;

const QUESTION_TYPES = {
    OPEN: 'open',
    FIVE_POINT_LIKERT_SCALE: 'five-point-likert-scale',
};

const FIVE_POINT_LIKERT_SCALE_RESPONSES = [
    {value: '1', text: 'Strongly Agree'},
    {value: '2', text: 'Agree'},
    {value: '3', text: 'Neutral'},
    {value: '4', text: 'Disagree'},
    {value: '5', text: 'Strongly Disagree'},
];

//
// Export the constants
//
export default {
    ACTION_TYPES,
    FIVE_POINT_LIKERT_SCALE_RESPONSES,
    OPEN_QUESTION_INITIAL_ROWS,
    OPEN_QUESTION_MAX_LENGTH,
    PLUGIN_NAME,
    QUESTION_TYPES,
};
