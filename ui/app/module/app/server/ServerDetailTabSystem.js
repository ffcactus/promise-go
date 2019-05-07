import React from "react";
import PropTypes from "prop-types";
import ServerDetailSectionProcessor from "./ServerDetailSectionProcessor";
import ServerDetailSectionMemory from "./ServerDetailSectionMemory";
import ServerDetailSectionEthernetInterface from "./ServerDetailSectionEthernetInterface";
import ServerDetailSectionNetworkInterface from "./ServerDetailSectionNetworkInterface";
import ServerDetailSectionStorage from "./ServerDetailSectionStorage";

function ServerDetailTabSystem(props) {
  let content;
  if (props.computerSystem === null) {
    content = <p>Empty</p>;
  } else {
    content = (
      <div style={{ height: "100%", overflow: "auto" }}>
        <ServerDetailSectionProcessor
          processors={props.computerSystem.Processors}
        />
        <hr />
        <ServerDetailSectionMemory memory={props.computerSystem.Memory} />
        <hr />
        <ServerDetailSectionNetworkInterface
          networkInterfaces={props.computerSystem.NetworkInterfaces}
        />
        <hr />
        <ServerDetailSectionStorage storages={props.computerSystem.Storages} />
        <hr />
        <ServerDetailSectionEthernetInterface
          ethernetInterfaces={props.computerSystem.EthernetInterfaces}
        />
      </div>
    );
  }
  return content;
}

ServerDetailTabSystem.propTypes = {
  computerSystem: PropTypes.object
};

export default ServerDetailTabSystem;
