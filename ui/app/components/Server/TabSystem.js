import React from 'react';
import PropTypes from 'prop-types';
import SectionProcessor from './SectionProcessor';
import SectionMemory from './SectionMemory';
import SectionEthernetInterface from './SectionEthernetInterface';
import SectionNetworkInterface from './SectionNetworkInterface';
import SectionStorage from './SectionStorage';

function TabSystem(props) {
  let content;
  if (props.computerSystem === null) {
    content = <p>Empty</p>;
  } else {
    content = (<div>
      <SectionProcessor processors={props.computerSystem.Processors} />
      <hr />
      <SectionMemory memory={props.computerSystem.Memory} />
      <hr />
      <SectionNetworkInterface networkInterfaces={props.computerSystem.NetworkInterfaces} />
      <hr />
      <SectionStorage storages={props.computerSystem.Storages} />
      <hr />
      <SectionEthernetInterface ethernetInterfaces={props.computerSystem.EthernetInterfaces} />
    </div>);
  }
  return content;
}

TabSystem.propTypes = {
  computerSystem: PropTypes.object
};

export default TabSystem;
