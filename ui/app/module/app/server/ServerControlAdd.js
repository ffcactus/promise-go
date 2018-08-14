import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import CenterDiv from '../../promise/common/CenterDiv';

/**
 * ServerControlAdd is the button to pop a dialog for server adding.
 */
class ServerControlAdd extends React.Component {
  constructor(props) {
    super(props);
    // this.onClick = this.onClick.bind(this);
  }

  render() {
    const icon = require('../../promise/common/img/icon/Navigation_Add_2x.png');
    return (
      <div styleName="ListControlAreaButton" style={{float: 'right'}}>
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

ServerControlAdd.propTypes = {
  dispatch: PropTypes.func,
  openCreateServerGroupDialog: PropTypes.bool,
};

export default connect(mapStateToProps)(CSSModules(ServerControlAdd, styles));
