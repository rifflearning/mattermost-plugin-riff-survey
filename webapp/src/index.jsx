import PostTypeSurvey from './components/post_type_survey';
import SurveyModal from './components/survey_modal';

import Actions from './actions';
import Constants from './constants';
import reducer from './reducers';

// Global Styles
import './styles.css';

//
// Define the plugin class that will register
// our plugin components.
//
class PluginClass {
    store;

    /**
     * initialize is called by the webapp when the plugin is first loaded.
     * Receives the following:
     * - registry - an instance of the registry tied to your plugin id
     * - store - the Redux store of the web app.
     */
    initialize(registry, store) {
        this.store = store;
        registry.registerReducer(reducer);
        store.dispatch(Actions.pluginEnabled());
        registry.registerRootComponent(SurveyModal);
        registry.registerPostTypeComponent(
            'custom_survey',
            PostTypeSurvey,
        );
    }

    /**
     * uninitialize is called by the webapp if your plugin is uninstalled
     */
    uninitialize() {
        this.store.dispatch(Actions.pluginDisabled());
    }
}

//
// To register your plugin you must expose it
// on window.
//
window.registerPlugin(Constants.PLUGIN_NAME, new PluginClass());
