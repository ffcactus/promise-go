import React from "react";
import CSSModules from "react-css-modules";
import ServerControlDiscover from "./ServerControlDiscover";
import ServerControlEdit from "./ServerControlEdit";
import styles from "./App.css";

class ServerDetailControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div
        styleName="flex-row-container border-column-first"
        style={{ height: "40px" }}
      >
        <ServerControlDiscover />
        <ServerControlEdit />
      </div>
    );
  }
}

export default CSSModules(ServerDetailControlArea, styles, {
  allowMultiple: true
});
