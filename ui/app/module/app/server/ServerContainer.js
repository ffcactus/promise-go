import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import AppInitFailure from '../../promise/common/AppInitFailure';
import Server from './Server';
import { AppState } from './ConstValue';
import * as AppAction from './AppAction';
import * as Util from '../../promise/common/Util';

class ServerContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    const search = this.props.location.search;
    const servergroup = Util.getParameterByName('servergroup', search);
    const server = Util.getParameterByName('server', search);
    this.props.dispatch(AppAction.appInit(servergroup, server));
  }

  render() {
    switch (this.props.appState) {
      case AppState.LOADING:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
      case AppState.NORMAL:
        return <Server />;
      case AppState.FAILURE:
        return <AppInitFailure />;
      default:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
    }
  }
}

function mapStateToProps(state) {
  return {
    appState: state.serverApp.appState
  };
}

ServerContainer.propTypes = {
  appState: PropTypes.string,
  location: PropTypes.object,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(ServerContainer);
