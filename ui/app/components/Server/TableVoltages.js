import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';

function TableVoltages(props) {
  return (<table>
    <thead>
      <tr>
        <th colSpan="8" styleName="level1">Voltage</th>
      </tr>
      <tr>
        <th styleName="level2">Name</th>
        <th styleName="level2">Value</th>
        <th styleName="level2">Upper Fatal</th>
        <th styleName="level2">Upper Critical</th>
        <th styleName="level2">Upper NonCritical</th>
        <th styleName="level2">Lower NonCritical</th>
        <th styleName="level2">Lower Critical</th>
        <th styleName="level2">Lower Fatal</th>
      </tr>
    </thead>
    <tbody>
      {
        props.voltages.map((each, i) => {
          return (
            <tr key={i.toString()}>
              <td>{each.Name}</td>
              <td>{each.ReadingVolts}</td>
              <td>{each.UpperThresholdFatal}</td>
              <td>{each.UpperThresholdCritical}</td>
              <td>{each.UpperThresholdNonCritical}</td>
              <td>{each.LowerThresholdNonCritical}</td>
              <td>{each.LowerThresholdCritical}</td>
              <td>{each.LowerThresholdFatal}</td>
            </tr>
          );
        })
      }
    </tbody>
  </table>);
}

TableVoltages.propTypes = {
  voltages: PropTypes.array
};

export default CSSModules(TableVoltages, styles);
