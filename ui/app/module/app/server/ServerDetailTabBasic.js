import React from "react";
import PropTypes from "prop-types";
import ServerDetailSectionBasic from "./ServerDetailSectionBasic";

function ServerDetailTabBasic(props) {
  let content;
  if (props.computerSystem === null) {
    content = <p>Empty</p>;
  } else {
    content = (
      <div style={{ height: "100%", overflow: "auto" }}>
        <ServerDetailSectionBasic server={props.server} />
      </div>
    );
  }
  return content;
}

ServerDetailTabBasic.propTypes = {
  server: PropTypes.object
};

export default ServerDetailTabBasic;
