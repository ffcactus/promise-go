import React from 'react';
import PropTypes from 'prop-types';
import SectionHeader from './SectionHeader';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';

function SectionProcessor(props) {
  const iconImage = require('../../img/icon/Processor.png');
  return (<div styleName="sectionDiv">
    <SectionHeader name="Processor" image={iconImage} />
    <table>
      <thead>
        <tr>
          <th styleName="level1">Name</th>
          <th styleName="level1">Model</th>
          <th styleName="level1">Architecture</th>
          <th styleName="level1">InstructionSet</th>
          <th styleName="level1">Socket</th>
          <th styleName="level1">MaxSpeedMHz</th>
          <th styleName="level1">TotalCores</th>
        </tr>
      </thead>
      <tbody>
        {
          props.processors.map(each => {
            return (
              <tr key={each.Name}>
                <td>{each.Name}</td>
                <td>{each.Model}</td>
                <td>{each.ProcessorArchitecture}</td>
                <td>{each.InstructionSet}</td>
                <td>{each.Socket}</td>
                <td>{each.MaxSpeedMHz}</td>
                <td>{each.TotalCores}</td>
              </tr>
            );
          })
        }
      </tbody>
    </table>
  </div>);
}

SectionProcessor.propTypes = {
  processors: PropTypes.array
};

export default CSSModules(SectionProcessor, styles);
