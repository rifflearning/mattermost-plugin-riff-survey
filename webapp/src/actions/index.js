import {
    closeSurveyModal, getSurvey, getSurveyResponses, openSurveyModal, submitSurveyResponses,
} from './survey';
import {setCurrentPostID, setCurrentPostProps} from './post';
import {openRiffDashboard} from './dashboard';
import {pluginEnabled, pluginDisabled} from './enable';

export default {
    closeSurveyModal,
    getSurvey,
    getSurveyResponses,
    openRiffDashboard,
    openSurveyModal,
    pluginEnabled,
    pluginDisabled,
    setCurrentPostID,
    setCurrentPostProps,
    submitSurveyResponses,
};
