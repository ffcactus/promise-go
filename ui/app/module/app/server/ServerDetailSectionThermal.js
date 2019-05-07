import React from "react";
import PropTypes from "prop-types";
import CSSModules from "react-css-modules";
import ServerDetailSectionHeader from "./ServerDetailSectionHeader";
import ServerDetailTableFan from "./ServerDetailTableFan";
import styles from "./App.css";

function ServerDetailSectionThermal(props) {
  const iconImage = require("./img/icon/Thermal.png");
  return (
    <div styleName="ServerDetailSectionDiv">
      <ServerDetailSectionHeader name="Thermal" image={iconImage} />
      <ServerDetailTableFan fan={props.thermal.Fans} />
    </div>
  );
}

ServerDetailSectionThermal.propTypes = {
  thermal: PropTypes.object
};

export default CSSModules(ServerDetailSectionThermal, styles);
