import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

class ServerGroupElement extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const currentStyle = this.props.server.currentServerGroup === this.props.name ? "ServerGroupElementSelected" : "ServerGroupElement";
    return (
      <div styleName={currentStyle}>
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
};

export default connect(mapStateToProps)(CSSModules(ServerGroupElement, styles));
