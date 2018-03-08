import React from 'react';
import PropTypes from 'prop-types';
import SectionHeader from './SectionHeader';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';

function SectionNetworkInterface(props) {
  const iconImage = require('../../img/icon/NetworkInterface.png');

  return (<div styleName="sectionDiv">
    <SectionHeader name="Network Interface" image={iconImage} />
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

SectionNetworkInterface.propTypes = {
  networkInterfaces: PropTypes.array,
};

export default CSSModules(SectionNetworkInterface, styles);
