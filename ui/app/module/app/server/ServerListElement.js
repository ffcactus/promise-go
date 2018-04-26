import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import * as ServerAction from './ServerAction';

class ServerListElement extends React.Component {
  constructor(props) {
    super(props);
    this.onSelect = this.onSelect.bind(this);
    this.state = {
      selected: false
    };
  }

  // on mounting we need send REST to get server to display some basic info.
  // TODO: should get collection operation return such info?
  componentDidMount() {
    this.props.dispatch(ServerAction.getServer(this.props.serverUri));
  }

  // On selecting we need display the detail infomation.
  onSelect(event) {
    event.preventDefault();
    this.props.dispatch(ServerAction.selectServer(this.props.serverUri));
  }

  render() {
    const currentStyle = 'ServerListElement ' + (
      this.props.serverApp.currentServer.URI === this.props.serverUri ?
        'Selected' : 'NotSelected'
    );
    const server = this.props.serverApp.serverList.get(this.props.serverUri);

    return (
      <div styleName={currentStyle} onClick={this.onSelect}>{server.Name}</div>
    );
  }
}

function mapStateToProps(state) {
  const { serverApp } = state;
  return { serverApp };
}

ServerListElement.propTypes = {
  serverApp: PropTypes.object,
  serverUri: PropTypes.string,
  children: PropTypes.string,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(CSSModules(ServerListElement, styles, {allowMultiple: true}));

