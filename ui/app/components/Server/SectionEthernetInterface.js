import React from 'react';
import PropTypes from 'prop-types';
import SectionHeader from './SectionHeader';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';
import ObjectEthernetInterface from './ObjectEthernetInterface';

function SectionEthernetInterface(props) {
  const iconImage = require('../../img/icon/EthernetInterface.png');
  return (<div styleName="sectionDiv">
    <SectionHeader name="Ethernet Interface" image={iconImage} />
    {
      props.ethernetInterfaces.map((each, i) => {
        return <ObjectEthernetInterface key={i.toString()} ethernetInterface={each} />;
      })
    }
  </div>);
}

SectionEthernetInterface.propTypes = {
  ethernetInterfaces: PropTypes.array
};

export default CSSModules(SectionEthernetInterface, styles);
