import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import ServerListElement from './ServerListElement';
import styles from './Server.css';
import { AutoSizer, List } from 'react-virtualized';
import 'react-virtualized/styles.css';

class ServerList extends React.Component {
  constructor(props) {
    super(props);
    this.rowRenderer = this.rowRenderer.bind(this);
  }

  rowRenderer({
    key,            // Unique key within array of rows
    index,          // Index of row within collection
    isScrolling,    // The List is currently being scrolled
    isVisible,      // This row is visible within the List
    style           // Style object to be applied to row
  }) {
    const server = this.props.serverList.get(index);
    return (
      <div key={key} style={style}>
        <ServerListElement server={server}/>
      </div>
    );
  }

  render() {
    return (
      <AutoSizer>{
        ({ height, width }) => (
          <List
            width={width}
            height={height}
            rowCount={this.props.serverList.size}
            rowHeight={40}
            rowRenderer={this.rowRenderer}
          />
        )
      }</AutoSizer>
    );
  }

  // render() {
  //   return (
  //     <div styleName="ServerList">
  //       <List
  //         width={452}
  //         height={867}
  //         rowCount={this.props.serverList.size}
  //         rowHeight={40}
  //         rowRenderer={this.rowRenderer}
  //       />
  //     </div>
  //   );
  // }
}

function mapStateToProps(state) {
  return {
    serverList: state.serverApp.serverList,
    currentServerSet: state.serverApp.currentServerSet,
  };
}

ServerList.propTypes = {
  serverList: PropTypes.object,
  currentServerSet: PropTypes.object,
  dispatch: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(ServerList, styles));

