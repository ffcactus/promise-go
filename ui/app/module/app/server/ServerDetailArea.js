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
  const server = props.serverApp.serverList.get(props.serverApp.currentServer);
  if (!server || !server.URI) {
    return <div />;
  }
  const pages = [
    {
      'title': 'Basic',
      'content': <ServerDetailTabBasic server={server} />
    },
    {
      'title': 'System',
      'content': <ServerDetailTabSystem computerSystem={server.ComputerSystem}/>
    },
    {
      'title': 'Chassis',
      'content': <ServerDetailTabChassis chassis={server.Chassis}/>
    }
  ];
  return (
    <Tab pages={pages} />
  );
}

ServerDetailArea.propTypes = {
  serverApp: PropTypes.object,
};

function mapStateToProps(state) {
  const { serverApp } = state;
  return { serverApp };
}

export default connect(mapStateToProps)(CSSModules(ServerDetailArea, styles, {allowMultiple: true}));
