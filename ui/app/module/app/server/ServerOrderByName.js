import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import CenterDiv from '../../promise/common/CenterDiv';
import * as ServerAction from './ServerAction';

class ServerOrderByName extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
    this.props.dispatch(ServerAction.sortByName());
  }

  render() {
    const icon = require('../../promise/common/img/icon/Navigation_Add_2x.png');
    return (
      <div styleName="ServerGroupControlButton">
        <CenterDiv><img src={icon} onClick={this.onClick} style={{display: 'block', margin: 'auto', height: '30px'}}/></CenterDiv>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    serverSortedBy: state.serverApp.serverSortedBy
  };
}

ServerOrderByName.propTypes = {
  dispatch: PropTypes.func,
  serverSortByName: PropTypes.bool,
};

export default connect(mapStateToProps)(CSSModules(ServerOrderByName, styles));
