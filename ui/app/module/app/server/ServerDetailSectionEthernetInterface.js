import React from 'react';
import PropTypes from 'prop-types';
import ServerDetailSectionHeader from './ServerDetailSectionHeader';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerDetailObjectEthernetInterface from './ServerDetailObjectEthernetInterface';

function ServerDetailSectionEthernetInterface(props) {
  const iconImage = require('./img/icon/EthernetInterface.png');
  return (<div styleName="ServerDetailSectionDiv">
    <ServerDetailSectionHeader name="Ethernet Interface" image={iconImage} />
    {
      props.ethernetInterfaces.map((each, i) => {
        return <ServerDetailObjectEthernetInterface key={i.toString()} ethernetInterface={each} />;
      })
    }
  </div>);
}

ServerDetailSectionEthernetInterface.propTypes = {
  ethernetInterfaces: PropTypes.array
};

export default CSSModules(ServerDetailSectionEthernetInterface, styles);
