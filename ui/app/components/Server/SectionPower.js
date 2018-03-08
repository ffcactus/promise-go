import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import SectionHeader from './SectionHeader';
import TableVoltages from './TableVoltages';
import styles from '../../styles/ServerFrame.css';

function SectionPower(props) {
  const iconImage = require('../../img/icon/Power.png');
  return (<div styleName="sectionDiv">
    <SectionHeader name="Power" image={iconImage} />
    <TableVoltages voltages={props.power.Voltages} />
    <table>
      <thead>
        <tr>
          <th styleName="level1">Name</th>
        </tr>
      </thead>
      <tbody>
      </tbody>
    </table>
  </div>);
}

SectionPower.propTypes = {
  power: PropTypes.object
};

export default CSSModules(SectionPower, styles);
