import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import Tab from '../../promise/common/Tab';
import TabSystem from './TabSystem';
import TabChassis from './TabChassis';
import TabBasic from './TabBasic';

function ServerDetailArea(props) {
  const server = props.serverApp.serverList.get(props.serverApp.currentServer);
  if (!server || !server.URI) {
    return <div />;
  }
  const pages = [
    {
      'title': 'Basic',
      'content': <TabBasic server={server} />
    },
    {
      'title': 'System',
      'content': <TabSystem computerSystem={server.ComputerSystem}/>
    },
    {
      'title': 'Chassis',
      'content': <TabChassis chassis={server.Chassis}/>
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
