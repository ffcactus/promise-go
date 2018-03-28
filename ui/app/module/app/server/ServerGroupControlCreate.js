import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
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
    if (! this.props.server.openCreateServerGroupDialog) {
      this.props.dispatch(ServerGroupAction.openCreateServerGroupDialog());
    }
  }

  render() {
    const icon = require('../../promise/common/img/icon/Navigation_Add_2x.png');
    return (
      <div styleName="ServerGroupControlButton">
        <CreateServerGroupDialog />
        <CenterDiv><img src={icon} onClick={this.onClick} style={{display: 'block', margin: 'auto', height: '30px'}}/></CenterDiv>
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { server } = state;
  return { server };
}

ServerGroupControlCreate.propTypes = {
  dispatch: PropTypes.func,
  server: PropTypes.object,
};

export default connect(mapStateToProps)(CSSModules(ServerGroupControlCreate, styles));

