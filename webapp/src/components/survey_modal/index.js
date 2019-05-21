import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import Actions from '../../actions';
import Selectors from '../../selectors';

import SurveyModal from './survey_modal';

const mapStateToProps = (state) => ({
    visible: Selectors.isSurveyModalVisible(state),
    currentPostProps: Selectors.currentPostProps(state),
});

const mapDispatchToProps = (dispatch) => bindActionCreators({
    close: Actions.closeSurveyModal,
    getSurvey: Actions.getSurvey,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(SurveyModal);
