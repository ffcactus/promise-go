import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import Server from './Server';
import { ServerAppState } from './ConstValue';
import * as Action from './ServerAction';

class ServerContainer extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      hostname: window.location.hostname
    };
  }

  componentWillMount() {
    this.props.dispatch(Action.appInit(this.state.hostname));
  }

  render() {
    switch (this.props.server.state) {
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
  const { server } = state;
  return { server };
}

ServerContainer.propTypes = {
  hostname: PropTypes.string,
  server: PropTypes.object,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(ServerContainer);
