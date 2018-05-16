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

  onClick(event) {
    event.preventDefault();
    this.props.dispatch(ServerGroupAction.uiListSelect(this.props.servergroup.URI));
  }

  render() {
    const currentStyle = 'ServerGroupElement ' + (this.props.currentServerGroupUri === this.props.servergroup.URI ? 'Selected' : 'NotSelected');
    return (
      <div styleName={currentStyle} onClick={this.onClick}>
        <p styleName="ServerGroupElementText">{this.props.servergroup.Name}</p>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return { currentServerGroupUri: state.serverApp.currentServerGroupUri };
}

ServerGroupElement.propTypes = {
  servergroup: PropTypes.object,
  dispatch: PropTypes.func,
  currentServerGroupUri: PropTypes.string,
};

export default connect(mapStateToProps)(CSSModules(ServerGroupElement, styles, {allowMultiple: true}));
