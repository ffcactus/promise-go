import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import ServerDetailSectionHeader from './ServerDetailSectionHeader';
import styles from './App.css';

function ServerDetailSectionMemory(props) {
  const iconImage = require('./img/icon/Memory.png');
  return (<div styleName="ServerDetailSectionDiv">
    <ServerDetailSectionHeader name="Memory" image={iconImage} />
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

ServerDetailSectionMemory.propTypes = {
  memory: PropTypes.array
};

export default CSSModules(ServerDetailSectionMemory, styles);
