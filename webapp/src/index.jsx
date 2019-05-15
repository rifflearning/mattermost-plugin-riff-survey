import PostTypeSurvey from './components/post_type_survey';
import SurveyModal from './components/survey_modal';

import reducer from './reducers';
import Constants from './constants';

// Global Styles
import './styles.css';

//
// Define the plugin class that will register
// our plugin components.
//
class PluginClass {
    initialize(registry) {
        registry.registerRootComponent(SurveyModal);
        registry.registerPostTypeComponent(
            'custom_survey',
            PostTypeSurvey,
        );
        registry.registerReducer(reducer);
    }
}

//
// To register your plugin you must expose it
// on window.
//
window.registerPlugin(Constants.PLUGIN_NAME, new PluginClass());
