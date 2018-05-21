import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import Server from './Server';
import { ServerAppState } from './ConstValue';
import * as ServerAppAction from './ServerAppAction';
import * as Util from '../../promise/common/Util';

class ServerContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    const search = this.props.location.search;
    const servergroup = Util.getParameterByName('servergroup', search);
    const server = Util.getParameterByName('server', search);
    this.props.dispatch(ServerAppAction.appInit(servergroup, server));
  }

  render() {
    switch (this.props.serverApp.appState) {
      case ServerAppState.LOADING:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
      case ServerAppState.NORMAL:
        return <Server />;
      case ServerAppState.FAILURE:
        return <CenterDiv><p>App initialization failure.</p></CenterDiv>;
      default:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
    }
  }
}

function mapStateToProps(state) {
  const { serverApp } = state;
  return { serverApp };
}

ServerContainer.propTypes = {
  serverApp: PropTypes.object,
  location: PropTypes.object,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(ServerContainer);
