import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import Actions from '../../actions';

import PostTypeSurvey from './post_type_survey';

const mapDispatchToProps = (dispatch) => bindActionCreators({
    open: Actions.openSurveyModal,
}, dispatch);

export default connect(null, mapDispatchToProps)(PostTypeSurvey);
