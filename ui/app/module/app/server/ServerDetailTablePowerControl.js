import React from "react";
import PropTypes from "prop-types";
import CSSModules from "react-css-modules";
import styles from "./App.css";

function ServerDetailTablePowerControl(props) {
  return (
    <table>
      <thead>
        <tr>
          <th colSpan="7" styleName="level1">
            PowerControl
          </th>
        </tr>
        <tr>
          <th styleName="level2">PowerConsumedWatts</th>
          <th styleName="level2">PowerCapacityWatts</th>
          <th styleName="level2">MinConsumedWatts</th>
          <th styleName="level2">MaxConsumedWatts</th>
          <th styleName="level2">AverageConsumedWatts</th>
          <th styleName="level2">LimitInWatts</th>
          <th styleName="level2">LimitException</th>
        </tr>
      </thead>
      <tbody>
        {props.powerControl.map((each, i) => {
          return (
            <tr key={i.toString()}>
              <td>{each.PowerConsumedWatts}</td>
              <td>{each.PowerCapacityWatts}</td>
              <td>{each.PowerMetrics.MinConsumedWatts}</td>
              <td>{each.PowerMetrics.MaxConsumedWatts}</td>
              <td>{each.PowerMetrics.AverageConsumedWatts}</td>
              <td>{each.PowerLimit.LimitInWatts}</td>
              <td>{each.PowerLimit.LimitException}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
}

ServerDetailTablePowerControl.propTypes = {
  powerControl: PropTypes.array
};

export default CSSModules(ServerDetailTablePowerControl, styles);
