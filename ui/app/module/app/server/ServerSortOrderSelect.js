import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './App.css';
import * as ServerAction from './ServerAction';

class ServerSortOrderSelect extends React.Component {
  constructor(props) {
    super(props);
    this._onChange = this._onChange.bind(this);
  }

  _onChange(event) {
    event.preventDefault();
    this.props.dispatch(ServerAction.onServerOrderChange(event.currentTarget.value));
    // this.props.listRef.scrollToRow(0);
    // this.props.listRef.forceUpdateGrid();
  }

  render() {
    return (
      <select onChange={this._onChange}>{
        this.props.options.map(each => {
          return <option key={each} value={each}>{each}</option>;
        })
      }</select>
    );
  }
}

function mapStateToProps(state) {
  return {
    serverSortOrder: state.serverApp.serverSortOrder
  };
}

ServerSortOrderSelect.propTypes = {
  // listRef: PropTypes.object,
  dispatch: PropTypes.func,
  serverSortOrder: PropTypes.string,
  options: PropTypes.arrayOf(PropTypes.string),
};

export default connect(mapStateToProps)(CSSModules(ServerSortOrderSelect, styles));
