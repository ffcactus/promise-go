import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerSearchArea from './ServerSearchArea';
import ServerListControlArea from './ServerListControlArea';
import ServerList from './ServerList';

class ServerListArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const size = this.props.serverList.size;
    const serverList = [];
    for (let i = 0; i < size; i++) {
      const server = this.props.serverList.get(i);
      if (this.props.currentServerSet.has(server.URI)) {
        serverList.push(server);
      }
    }
    return (
      <div styleName="ServerListArea">
        <ServerSearchArea />
        <ServerListControlArea />
        <div style={{ flex: '1 1 auto' }}>
          <ServerList serverList={serverList} />
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    serverList: state.serverApp.serverList,
    currentServerSet: state.serverApp.currentServerSet,
  };
}

ServerListArea.propTypes = {
  serverList: PropTypes.object,
  currentServerSet: PropTypes.object,
  dispatch: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(ServerListArea, styles));

