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
    if (this.props.serverApp.currentServerGroup.ID === this.props.servergroup.ID) {
      this.props.dispatch(ServerGroupAction.selectServerGroup(this.props.servergroup));
    }
  }

  onClick(event) {
    event.preventDefault();
    this.props.dispatch(ServerGroupAction.selectServerGroup(this.props.servergroup));
  }

  render() {
    const currentStyle = 'ServerGroupElement ' + (this.props.serverApp.currentServerGroup.ID === this.props.servergroup.ID ? 'Selected' : 'NotSelected');
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
