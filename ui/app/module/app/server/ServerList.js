import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import * as ServerAction from './ServerAction';
import ServerListElement from './ServerListElement';
import styles from './Server.css';

class ServerList extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    // At this point, the servergroup hasn't been clicked yet, so let's do a auto get collection.
    this.props.dispatch(ServerAction.onListDidMount());
  }

  render() {
    const size = this.props.serverList.size;
    if (size === 0) {
      return <div styleName="ServerList"><p>Empty</p></div>;
    }
    const list = [];
    for (let i = 0; i < size; i++) {
      const server = this.props.serverList.get(i);
      list.push(<ServerListElement key={server.key} server = {server}/>);
    }
    return <div styleName="ServerList">{list}</div>;
  }
}

function mapStateToProps(state) {
  return { serverList: state.serverApp.serverList};
}

ServerList.propTypes = {
  serverList: PropTypes.object,
  dispatch: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(ServerList, styles));

