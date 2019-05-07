import React from "react";
import PropTypes from "prop-types";
import ServerDetailSectionPower from "./ServerDetailSectionPower";
import ServerDetailSectionThermal from "./ServerDetailSectionThermal";

function ServerDetailTabChassis(props) {
  let content = null;
  if (props.chassis === null) {
    content = <p>Empty</p>;
  } else {
    content = (
      <div style={{ height: "100%", overflow: "auto" }}>
        <ServerDetailSectionPower power={props.chassis.Power} />
        <ServerDetailSectionThermal thermal={props.chassis.Thermal} />
        <hr />
      </div>
    );
  }
  return content;
}

ServerDetailTabChassis.propTypes = {
  chassis: PropTypes.object
};

export default ServerDetailTabChassis;
