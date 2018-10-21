import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './App.css';
import CenterDiv from '../../promise/common/CenterDiv';
import CreateServerGroupDialog from './CreateServerGroupDialog';
import * as ServerGroupAction from './ServerGroupAction';


class ServerGroupControlCreate extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
    if (! this.props.openCreateServerGroupDialog) {
      this.props.dispatch(ServerGroupAction.openCreateServerGroupDialog());
    }
  }

  render() {
    const icon = require('../../promise/common/img/icon/Navigation_Add_2x.png');
    return (
      <div styleName="ListControlAreaButton" style={{float: 'left'}}>
        <CreateServerGroupDialog />
        <CenterDiv><img src={icon} onClick={this.onClick} style={{display: 'block', margin: 'auto', height: '30px'}}/></CenterDiv>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    openCreateServerGroupDialog: state.serverApp.openCreateServerGroupDialog
  };
}

ServerGroupControlCreate.propTypes = {
  dispatch: PropTypes.func,
  openCreateServerGroupDialog: PropTypes.bool,
};

export default connect(mapStateToProps)(CSSModules(ServerGroupControlCreate, styles));

