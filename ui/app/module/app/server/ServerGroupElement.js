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
    this.props.dispatch(ServerGroupAction.onServerGroupSelected(this.props.name));
  }

  render() {
    const currentStyle = this.props.server.currentServerGroup === this.props.name ? 'ServerGroupElementSelected' : 'ServerGroupElement';
    return (
      <div styleName={currentStyle} onClick={this.onClick}>
        <p styleName="ServerGroupElementText">{this.props.name}</p>
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { server } = state;
  return { server };
}

ServerGroupElement.propTypes = {
  name: PropTypes.string,
  dispatch: PropTypes.func,
  server: PropTypes.object,
};

export default connect(mapStateToProps)(CSSModules(ServerGroupElement, styles));
