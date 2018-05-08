import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import SectionHeader from './SectionHeader';
import TableVoltages from './TableVoltages';
import styles from './Server.css';

function SectionPower(props) {
  const iconImage = require('./img/icon/Power.png');
  return (<div styleName="ServerDetailSectionDiv">
    <SectionHeader name="Power" image={iconImage} />
    <TableVoltages voltages={props.power.Voltages} />
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

SectionPower.propTypes = {
  power: PropTypes.object
};

export default CSSModules(SectionPower, styles);
