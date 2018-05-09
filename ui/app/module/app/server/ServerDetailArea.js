import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import Tab from '../../promise/common/Tab';
import ServerDetailTabSystem from './ServerDetailTabSystem';
import ServerDetailTabChassis from './ServerDetailTabChassis';
import ServerDetailTabBasic from './ServerDetailTabBasic';

function ServerDetailArea(props) {
  if (!props.currentServer) {
    return <div />;
  }
  const pages = [
    {
      'title': 'Basic',
      'content': <ServerDetailTabBasic server={props.currentServer} />
    },
    {
      'title': 'System',
      'content': <ServerDetailTabSystem computerSystem={props.currentServer.ComputerSystem}/>
    },
    {
      'title': 'Chassis',
      'content': <ServerDetailTabChassis chassis={props.currentServer.Chassis}/>
    }
  ];
  return (
    <Tab pages={pages} />
  );
}

ServerDetailArea.propTypes = {
  currentServer: PropTypes.object,
};

function mapStateToProps(state) {
  return { currentServer: state.serverApp.currentServer };
}

export default connect(mapStateToProps)(CSSModules(ServerDetailArea, styles, {allowMultiple: true}));
