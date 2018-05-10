import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import * as ServerGroupAction from './ServerGroupAction';
import styles from './Server.css';

class ServerGroupElement extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  componentDidMount() {
    if (this.props.serverApp.currentServerGroup && this.props.serverApp.currentServerGroupUri === this.props.servergroup.URI) {
      this.props.dispatch(ServerGroupAction.uiListSelect(this.props.servergroup.URI));
    }
  }

  onClick(event) {
    event.preventDefault();
    this.props.dispatch(ServerGroupAction.uiListSelect(this.props.servergroup.URI));
  }

  render() {
    const currentStyle = 'ServerGroupElement ' + (this.props.serverApp.currentServerGroupUri === this.props.servergroup.URI ? 'Selected' : 'NotSelected');
    return (
      <div styleName={currentStyle} onClick={this.onClick}>
        <p styleName="ServerGroupElementText">{this.props.servergroup.Name}</p>
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { serverApp } = state;
  return { serverApp };
}

ServerGroupElement.propTypes = {
  servergroup: PropTypes.object,
  dispatch: PropTypes.func,
  serverApp: PropTypes.object,
};

export default connect(mapStateToProps)(CSSModules(ServerGroupElement, styles, {allowMultiple: true}));
