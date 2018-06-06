import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import ServerProfile from './ServerProfile';
import { AppState } from './ConstValue';
import * as AppAction from './AppAction';
import * as Util from '../../promise/common/Util';

class ServerProfileContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  compnentDidMount() {
    const search = this.props.location.search;
    const model = Util.getParameterByName('model', search);
    const config = Util.getParameterByName('config', search);
    this.props.dispatch(AppAction.appInit(model, config));
  }

  render() {
    switch (this.props.appState) {
      case AppState.LOADING:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
      case AppState.NORMAL:
        return <ServerProfile />;
      case AppState.FAILURE:
        return <CenterDiv><p>App initialization failure.</p></CenterDiv>;
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

ServerProfileContainer.propTypes = {
  dispatch: PropTypes.func,
  location: PropTypes.object,
  appState: PropTypes.string
};

export default connect(mapStateToProps)(ServerProfileContainer);
