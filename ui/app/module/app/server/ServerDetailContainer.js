import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as ServerAction from './ServerAction';
import { ServerDetailState } from './ConstValue';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import ServerDetail from './ServerDetail';

class ServerDetailContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    if (this.props.currentServer === null && this.props.currentServerUri !== null) {
      this.props.dispatch(ServerAction.getServer(this.props.currentServerUri));
    }
  }

  render() {
    switch(this.props.serverDetailState) {
      case ServerDetailState.EMPTY:
        return <CenterDiv><p>No Server Selected.</p></CenterDiv>;
      case ServerDetailState.LOADING:
        return <CenterDiv><LoadingIcon/></CenterDiv>;
      case ServerDetailState.READY:
        return <ServerDetail server={this.props.currentServer} />;
      case ServerDetailState.FAILURE:
        return <CenterDiv><p>Loading Server Failed</p></CenterDiv>;
      default:
        return <CenterDiv><LoadingIcon/></CenterDiv>;
    }
  }
}

ServerDetailContainer.propTypes = {
  currentServer: PropTypes.object,
  currentServerUri: PropTypes.string,
  serverDetailState: PropTypes.string,
  dispatch: PropTypes.func,
};

function mapStateToProps(state) {
  return {
    currentServerUri: state.serverApp.currentServerUri,
    currentServer: state.serverApp.currentServer,
    serverDetailState: state.serverApp.serverDetailState,
  };
}

export default connect(mapStateToProps)(ServerDetailContainer);
