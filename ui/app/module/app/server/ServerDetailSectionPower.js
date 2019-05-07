import React from "react";
import PropTypes from "prop-types";
import CSSModules from "react-css-modules";
import ServerDetailSectionHeader from "./ServerDetailSectionHeader";
import ServerDetailTablePowerSupply from "./ServerDetailTablePowerSupply";
import ServerDetailTablePowerControl from "./ServerDetailTablePowerControl";
import styles from "./App.css";

function ServerDetailSectionPower(props) {
  const iconImage = require("./img/icon/Power.png");
  return (
    <div styleName="ServerDetailSectionDiv">
      <ServerDetailSectionHeader name="Power" image={iconImage} />
      <ServerDetailTablePowerSupply powerSupply={props.power.PowerSupplies} />
      <ServerDetailTablePowerControl powerControl={props.power.PowerControl} />
    </div>
  );
}

ServerDetailSectionPower.propTypes = {
  power: PropTypes.object
};

export default CSSModules(ServerDetailSectionPower, styles);
