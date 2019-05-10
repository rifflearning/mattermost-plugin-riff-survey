import PostTypeSurvey from './components/post_type_survey';

import Constants from './constants';

//
// Define the plugin class that will register
// our plugin components.
//
class PluginClass {
    initialize(registry) {
        registry.registerPostTypeComponent(
            'custom_survey',
            PostTypeSurvey,
        );
    }
}

//
// To register your plugin you must expose it
// on window.
//
window.registerPlugin(Constants.PLUGIN_NAME, new PluginClass());
