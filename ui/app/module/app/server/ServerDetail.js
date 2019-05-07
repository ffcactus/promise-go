import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import Tab from "../../promise/common/Tab";
import ServerDetailTabSystem from "./ServerDetailTabSystem";
import ServerDetailTabChassis from "./ServerDetailTabChassis";
import ServerDetailTabBasic from "./ServerDetailTabBasic";
import { ActionType, ServerTabState } from "./ConstValue";
import CSSModules from "react-css-modules";
import styles from "./App.css";

class ServerDetail extends React.Component {
  constructor(props) {
    super(props);
    this.onBasic = this.onBasic.bind(this);
    this.onChassis = this.onChassis.bind(this);
    this.onSystem = this.onSystem.bind(this);
  }

  onBasic() {
    this.props.dispatch({
      type: ActionType.SERVER_UI_TAB_CHANGE,
      info: ServerTabState.BASIC
    });
  }

  onChassis() {
    this.props.dispatch({
      type: ActionType.SERVER_UI_TAB_CHANGE,
      info: ServerTabState.CHASSIS
    });
  }

  onSystem() {
    this.props.dispatch({
      type: ActionType.SERVER_UI_TAB_CHANGE,
      info: ServerTabState.SYSTEM
    });
  }

  render() {
    if (!this.props.server) {
      return <div />;
    }
    let tabContent;
    if (this.props.currentServerTab === ServerTabState.BASIC) {
      tabContent = <ServerDetailTabBasic server={this.props.server} />;
    }
    if (this.props.currentServerTab === ServerTabState.CHASSIS) {
      tabContent = (
        <ServerDetailTabChassis chassis={this.props.server.Chassis} />
      );
    }
    if (this.props.currentServerTab === ServerTabState.SYSTEM) {
      tabContent = (
        <ServerDetailTabSystem
          computerSystem={this.props.server.ComputerSystem}
        />
      );
    }
    return (
      <div styleName="TabRoot">
        <div styleName="TabHeader">
          <button onClick={this.onBasic} style={{ height: "40px" }}>
            Basic
          </button>
          <button onClick={this.onChassis} style={{ height: "40px" }}>
            Chassis
          </button>
          <button onClick={this.onSystem} style={{ height: "40px" }}>
            System
          </button>
          <button style={{ height: "40px" }}>Control</button>
        </div>
        <div styleName="TabContent">{tabContent}</div>
      </div>
    );
  }
}

ServerDetail.propTypes = {
  server: PropTypes.object,
  currentServerTab: PropTypes.string,
  dispatch: PropTypes.func
};

function mapStateToProps(state) {
  return {
    currentServerTab: state.serverApp.currentServerTab
  };
}

export default connect(mapStateToProps)(
  CSSModules(ServerDetail, styles, { allowMultiple: true })
);
