import React from "react";
import PropTypes from "prop-types";
import CSSModules from "react-css-modules";
import styles from "./App.css";

function ServerDetailTableFan(props) {
  return (
    <table>
      <thead>
        <tr>
          <th colSpan="3" styleName="level1">
            Fan
          </th>
        </tr>
        <tr>
          <th styleName="level2">State</th>
          <th styleName="level2">Health</th>
          <th styleName="level2">Reading</th>
        </tr>
      </thead>
      <tbody>
        {props.fan.map((each, i) => {
          return (
            <tr key={i.toString()}>
              <td>{each.PhysicalState}</td>
              <td>{each.PhysicalHealth}</td>
              <td>{each.Reading + " " + each.ReadingUnits}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
}

ServerDetailTableFan.propTypes = {
  fan: PropTypes.array
};

export default CSSModules(ServerDetailTableFan, styles);
