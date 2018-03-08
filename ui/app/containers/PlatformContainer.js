import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CenterDiv from '../components/common/CenterDiv';
import LoadingIcon from '../components/common/LoadingIcon';
import ConfirmDialog from '../components/common/ConfirmDialog';
import Desktop from '../components/Desktop/Desktop';
import * as platformReducer from '../reducer/PlatformReducer';

const PLATFORM_STARTING = 'PLATFORM_STARTING';
const PLATFORM_READY = 'PLATFORM_READY';
const PLATFORM_CRITICAL = 'PLATFORM_CRITICAL';
const PLATFORM_UNKNOWN = 'PLATFORM_UNKNOWN';

class PlatformContainer extends React.Component {
  constructor(props) {
    super(props);
    this.getPlatformState = this.getPlatformState.bind(this);
    this.getLoadingIcon = this.getLoadingIcon.bind(this);
  }

  getPlatformState() {
    const platform = this.props.platform;
    const services = [platform.serverServiceState, platform.taskServiceState];
    let isStarting = true;
    let isReady = true;
    let isCritical = false;
    for (const each of services) {
      if (each !== platformReducer.SVC_STATE_STARTING) {
        isStarting = false;
      }
    }
    if (isStarting) {
      return PLATFORM_STARTING;
    }
    for (const each of services) {
      if (each !== platformReducer.SVC_STATE_READY) {
        isReady = false;
      }
    }
    if (isReady) {
      return PLATFORM_READY;
    }
    for (const each of services) {
      if (each === platformReducer.SVC_STATE_CRITICAL) {
        isCritical = true;
      }
    }
    if (isCritical) {
      return PLATFORM_CRITICAL;
    }
    return PLATFORM_UNKNOWN;
  }

  getLoadingIcon() {
    return <CenterDiv><LoadingIcon /></CenterDiv>;
  }

  render() {
    switch (this.getPlatformState()) {
      case PLATFORM_STARTING:
      case PLATFORM_UNKNOWN:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
      case PLATFORM_READY:
        return <Desktop />;
      case PLATFORM_CRITICAL:
        return (
          <CenterDiv>
            <ConfirmDialog title="Platform Error" message="The platform turn to critical state, please login again."/>
          </CenterDiv>
        );
      default:
        return <p>Unknown</p>;
    }
  }
}

function mapStateToProps(state) {
  const { platform } = state;
  return { platform };
}

PlatformContainer.propTypes = {
  platform: PropTypes.object
};

export default connect(mapStateToProps)(PlatformContainer);
