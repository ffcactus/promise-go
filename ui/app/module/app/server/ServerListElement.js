import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import * as ServerAction from './ServerAction';
import ServerListElementTask from './ServerListElementTask';
import ServerListElementName from './ServerListElementName';

class ServerListElement extends React.Component {
  constructor(props) {
    super(props);
    this.onSelect = this.onSelect.bind(this);
    this.state = {
      selected: false
    };
  }

  // on mounting we need send REST to get server to display some basic info.
  componentDidMount() {
    this.props.dispatch(ServerAction.onElementDidMount(this.props.server.key));
  }

  // On selecting we need display the detail infomation.
  onSelect(event) {
    event.preventDefault();
    this.props.dispatch(ServerAction.uiListSelect(this.props.server.key));
  }

  render() {
    const currentStyle = 'ServerListElement ' + (
      this.props.currentServerUri === this.props.server.key ?
        'Selected' : 'NotSelected'
    );
    return (
      <div styleName={currentStyle} onClick={this.onSelect}>
        <ServerListElementName name={this.props.server.value ? this.props.server.value.Name : '...'} />
        <ServerListElementTask serverUri={this.props.server.key}/>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    // serverList: state.serverApp.serverList,
    currentServerUri: state.serverApp.currentServerUri,
  };
}

ServerListElement.propTypes = {
  server: PropTypes.object,
  currentServerUri: PropTypes.string,
  // serverList: PropTypes.object,
  // serverUri: PropTypes.string,
  children: PropTypes.string,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(CSSModules(ServerListElement, styles, {allowMultiple: true}));

