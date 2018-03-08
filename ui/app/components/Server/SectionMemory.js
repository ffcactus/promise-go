import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import SectionHeader from './SectionHeader';
import styles from '../../styles/ServerFrame.css';

function SectionMemory(props) {
  const iconImage = require('../../img/icon/Memory.png');
  return (<div styleName="sectionDiv">
    <SectionHeader name="Memory" image={iconImage} />
    <table>
      <thead>
        <tr>
          <th styleName="level1">Name</th>
          <th styleName="level1">Type</th>
          <th styleName="level1">Capacity (MB)</th>
          <th styleName="level1">Speed</th>
          <th styleName="level1">Manufacturer</th>
        </tr>
      </thead>
      <tbody>
        {
          props.memory.map(each => {
            return (
              <tr key={each.Name}>
                <td>{each.Name}</td>
                <td>{each.MemoryDeviceType}</td>
                <td>{each.CapacityMiB}</td>
                <td>{each.OperatingSpeedMhz}</td>
                <td>{each.Manufacturer}</td>
              </tr>
            );
          })
        }
      </tbody>
    </table>
  </div>);
}

SectionMemory.propTypes = {
  memory: PropTypes.array
};

export default CSSModules(SectionMemory, styles);
