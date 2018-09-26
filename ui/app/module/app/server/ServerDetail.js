import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Tab from '../../promise/common/Tab';
import ServerDetailTabSystem from './ServerDetailTabSystem';
import ServerDetailTabChassis from './ServerDetailTabChassis';
import ServerDetailTabBasic from './ServerDetailTabBasic';
import { ActionType } from './ConstValue';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

const ServerDetail = props => {
  if (!props.server) {
    return <div />;
  }
  const pages = [
    {
      'title': 'Basic',
      'content': <ServerDetailTabBasic server={props.server} />
    },
    {
      'title': 'System',
      'content': <ServerDetailTabSystem computerSystem={props.server.ComputerSystem}/>
    },
    {
      'title': 'Chassis',
      'content': <ServerDetailTabChassis chassis={props.server.Chassis}/>
    }
  ];
  const handler = (event) => {
    props.dispatch({
      type: ActionType.SERVER_UI_TAB_CHANGE,
      info: event,
    });
  };

  return (
    <div styleName="flex-item border-column" style={{height: '100%'}}>
      <Tab pages={pages} handler={handler} defaultOpen={props.currentServerTab} />
    </div>
  );
};

ServerDetail.propTypes = {
  server: PropTypes.object,
  currentServerTab: PropTypes.string,
  dispatch: PropTypes.func,
};

function mapStateToProps(state) {
  return {
    currentServerTab: state.serverApp.currentServerTab,
  };
}

export default connect(mapStateToProps)(CSSModules(ServerDetail, styles, {allowMultiple: true}));
