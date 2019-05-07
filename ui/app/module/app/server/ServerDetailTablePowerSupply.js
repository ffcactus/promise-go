import React from "react";
import PropTypes from "prop-types";
import CSSModules from "react-css-modules";
import styles from "./App.css";

function ServerDetailTablePowerSupply(props) {
  return (
    <table>
      <thead>
        <tr>
          <th colSpan="9" styleName="level1">
            PowerSupply
          </th>
        </tr>
        <tr>
          <th styleName="level2">State</th>
          <th styleName="level2">Health</th>
          <th styleName="level2">Model</th>
          <th styleName="level2">Manufacturer</th>
          <th styleName="level2">SerialNumber</th>
          <th styleName="level2">SparePartNumber</th>
          <th styleName="level2">PowerSupplyType</th>
          <th styleName="level2">PowerCapacityWatts</th>
          <th styleName="level2">FirmwareVersion</th>
        </tr>
      </thead>
      <tbody>
        {props.powerSupply.map((each, i) => {
          return (
            <tr key={i.toString()}>
              <td>{each.PhysicalState}</td>
              <td>{each.PhysicalHealth}</td>
              <td>{each.Model}</td>
              <td>{each.Manufacturer}</td>
              <td>{each.SerialNumber}</td>
              <td>{each.SparePartNumber}</td>
              <td>{each.PowerSupplyType}</td>
              <td>{each.PowerCapacityWatts}</td>
              <td>{each.FirmwareVersion}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
}

ServerDetailTablePowerSupply.propTypes = {
  powerSupply: PropTypes.array
};

export default CSSModules(ServerDetailTablePowerSupply, styles);
