import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import Server from './Server';
import { ServerAppState } from './ConstValue';
import * as ServerAppAction from './ServerAppAction';

class ServerContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  componentWillMount() {
    this.props.dispatch(ServerAppAction.appInit());
  }

  render() {
    switch (this.props.serverApp.state) {
      case ServerAppState.APP_INIT_START:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
      case ServerAppState.APP_INIT_SUCCESS:
        return <Server />;
      case ServerAppState.APP_INIT_FAILURE:
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
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(ServerContainer);
