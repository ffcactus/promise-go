import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as DesktopState from '../../constValue/state/Desktop';
import * as DesktopAction from '../../actions/DesktopAction';
import Desktop from '../../components/Desktop/Desktop';
import Login from '../../module/promise/login/Login';
import ServiceUnavaliable from '../../components/Desktop/ServiceUnavaliable';
import Background from '../../components/Desktop/Background';
import CenterDiv from '../../components/common/CenterDiv';
import LoadingIcon from '../../components/common/LoadingIcon';


class DesktopContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  componentWillMount() {
    this.props.dispatch(DesktopAction.init());
  }

  render() {
    let c = null;
    switch (this.props.desktop.state) {
      case DesktopState.GET_GLOBAL_START:
        c = <CenterDiv><LoadingIcon/></CenterDiv>;
        break;
      case DesktopState.GET_GLOBAL_SUCCESS:
        c = <Login/>;
        break;
      case DesktopState.GET_GLOBAL_FAILURE:
        c = <ServiceUnavaliable/>;
        break;
      default:
        c = <div><h1>Unknown</h1></div>;
        break;
    }
    return (
      <Desktop>
        <Background />
        {c}
      </Desktop>
    );
  }
}

function mapStateToProps(state) {
  const { desktop } = state;
  return { desktop };
}

DesktopContainer.propTypes = {
  desktop: PropTypes.object,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(DesktopContainer);
