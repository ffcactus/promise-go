import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import DiscoverServerDialog from './DiscoverServerDialog';
import * as ServerAction from './ServerAction';
import CenterDiv from '../../promise/common/CenterDiv';

/**
 * ServerControlAdd is the button to pop a dialog for server adding.
 */
class ServerControlDiscover extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
    if (! this.props.openCreateServerGroupDialog) {
      this.props.dispatch(ServerAction.openDiscoverServerDialog());
    }
  }

  render() {
    const icon = require('../../promise/common/img/icon/Navigation_Add_2x.png');
    return (
      <div styleName="ListControlAreaButton" style={{float: 'right'}}>
        <DiscoverServerDialog />
        <CenterDiv><img src={icon} onClick={this.onClick} style={{display: 'block', margin: 'auto', height: '30px'}}/></CenterDiv>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    openAddServerDialog: state.serverApp.openAddServerDialog
  };
}

ServerControlDiscover.propTypes = {
  dispatch: PropTypes.func,
  openCreateServerGroupDialog: PropTypes.bool,
};

export default connect(mapStateToProps)(CSSModules(ServerControlDiscover, styles));
