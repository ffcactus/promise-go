import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';

function ObjectStorage(props) {
  const o = props.storage;
  const controllers = o.StorageControllers;

  return (
    <table>
      <tbody>
        <tr>
          <th styleName="level1">Name</th>
          <td colSpan="4">{o.Name}</td>
        </tr>
        <tr>
          <th styleName="level1" rowSpan={(controllers.length + 1).toString()}>Controller</th>
          <th styleName="level2">Name</th>
          <th styleName="level2">Speed (GB)</th>
          <th styleName="level2">Firmware</th>
          <th styleName="level2">Protocol</th>
        </tr>
        {controllers.map(each => {
          return (
            <tr key={each.Name}>
              <td>{each.Name}</td>
              <td>{each.SpeedGbps}</td>
              <td>{each.FirmwareVersion}</td>
              <td>{JSON.stringify(each.SupportedDeviceProtocols)}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
}

ObjectStorage.propTypes = {
  storage: PropTypes.object,
};

export default CSSModules(ObjectStorage, styles);
