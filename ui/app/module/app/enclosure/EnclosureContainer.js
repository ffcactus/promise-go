import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import AppInitFailure from '../../promise/common/AppInitFailure';
import Enclosure from './Enclosure';
import { AppState } from './ConstValue';
import * as AppAction from './AppAction';

class EnclosureContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    this.props.dispatch(AppAction.appInit());
  }

  render() {
    switch (this.props.appState) {
      case AppState.LOADING:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
      case AppState.NORMAL:
        return <Enclosure />;
      case AppState.FAILURE:
        return <AppInitFailure />;
      default:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
    }
  }
}

function mapStateToProps(state) {
  return {
    appState: state.enclosureApp.appState
  };
}

EnclosureContainer.propTypes = {
  appState: PropTypes.string,
  location: PropTypes.object,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(EnclosureContainer);
