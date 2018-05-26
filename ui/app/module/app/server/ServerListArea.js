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
    this.listRef = React.createRef();
    this.setListRef = this.setListRef.bind(this);
  }

  setListRef(element) {
    this.listRef = element;
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
        <ServerListControlArea listRef={this.listRef}/>
        <div style={{ flex: '1 1 auto' }}>
          <ServerList setListRef={this.setListRef} serverList={serverList} serverIndex={this.props.serverIndex}/>
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    serverList: state.serverApp.serverList,
    serverIndex: state.serverApp.serverIndex,
    currentServerSet: state.serverApp.currentServerSet,
  };
}

ServerListArea.propTypes = {
  serverList: PropTypes.object,
  serverIndex: PropTypes.number,
  currentServerSet: PropTypes.object,
  dispatch: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(ServerListArea, styles));

