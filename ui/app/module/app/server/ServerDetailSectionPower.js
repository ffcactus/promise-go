import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import ServerDetailSectionHeader from './ServerDetailSectionHeader';
import ServerDetailTableVoltages from './ServerDetailTableVoltages';
import styles from './App.css';

function ServerDetailSectionPower(props) {
  const iconImage = require('./img/icon/Power.png');
  return (<div styleName="ServerDetailSectionDiv">
    <ServerDetailSectionHeader name="Power" image={iconImage} />
    <ServerDetailTableVoltages voltages={props.power.Voltages} />
    <table>
      <thead>
        <tr>
          <th styleName="level1">Name</th>
        </tr>
      </thead>
      <tbody/>
    </table>
  </div>);
}

ServerDetailSectionPower.propTypes = {
  power: PropTypes.object
};

export default CSSModules(ServerDetailSectionPower, styles);
