import React from 'react';
import PropTypes from 'prop-types';
import Tab from '../../promise/common/Tab';
import ServerDetailTabSystem from './ServerDetailTabSystem';
import ServerDetailTabChassis from './ServerDetailTabChassis';
import ServerDetailTabBasic from './ServerDetailTabBasic';

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
  return (
    <Tab pages={pages} />
  );
};

ServerDetail.propTypes = {
  server: PropTypes.object,
};

export default ServerDetail;
