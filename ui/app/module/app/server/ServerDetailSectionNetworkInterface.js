import React from 'react';
import PropTypes from 'prop-types';
import ServerDetailSectionHeader from './ServerDetailSectionHeader';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

function ServerDetailSectionNetworkInterface(props) {
  const iconImage = require('./img/icon/NetworkInterface.png');
  return (<div styleName="ServerDetailSectionDiv">
    <ServerDetailSectionHeader name="Network Interface" image={iconImage} />
    <table>
      <tbody>
        <tr>
          <th styleName="level1">Name</th>
          <th styleName="level1">Adapter</th>
        </tr>
        {props.networkInterfaces.map(each => {
          return (
            <tr key={each.Name}>
              <td>{each.Name}</td>
              <td>{each.NetworkAdapter.$ref}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  </div>);
}

ServerDetailSectionNetworkInterface.propTypes = {
  networkInterfaces: PropTypes.array,
};

export default CSSModules(ServerDetailSectionNetworkInterface, styles);
