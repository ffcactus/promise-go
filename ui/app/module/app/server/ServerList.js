import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import ServerListElement from './ServerListElement';
import styles from './App.css';
import { AutoSizer, List } from 'react-virtualized';

class ServerList extends React.Component {
  constructor(props) {
    super(props);
    this.rowRenderer = this.rowRenderer.bind(this);
  }

  rowRenderer({
    key,            // Unique key within array of rows
    index,          // Index of row within collection
    style           // Style object to be applied to row
  }) {
    const server = this.props.serverList[index];
    return (
      <div key={key} style={style}>
        <ServerListElement key={key} server={server} />
      </div>
    );
  }

  render() {
    return (
      <AutoSizer>{
        ({ height, width }) => (
          <List
            ref={this.props.setListRef}
            width={width}
            height={height}
            rowCount={this.props.serverList.length}
            scrollToIndex={this.props.serverIndex}
            rowHeight={40}
            rowRenderer={this.rowRenderer}
          />
        )
      }</AutoSizer>
    );
  }
}

ServerList.propTypes = {
  setListRef: PropTypes.func,
  serverIndex: PropTypes.number,
  serverList: PropTypes.array,
};

export default CSSModules(ServerList, styles);

