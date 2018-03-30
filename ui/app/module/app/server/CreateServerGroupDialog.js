import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import * as ServerGroupAction from './ServerGroupAction';
import DialogFrame from '../../promise/common/dialog/DialogFrame';
import DialogTitle from '../../promise/common/dialog/DialogTitle';
import DialogHR from '../../promise/common/dialog/DialogHR';
import DialogContentDiv from '../../promise/common/dialog/DialogContentDiv';
import DialogControlDiv from '../../promise/common/dialog/DialogControlDiv';
import DialogButton from '../../promise/common/dialog/DialogButton';


class CreateServerGroupDialog extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      discription: ''
    };
    this.handleNameChange = this.handleNameChange.bind(this);
    this.handleDiscriptionChange = this.handleDiscriptionChange.bind(this);
    this.handleOK = this.handleOK.bind(this);
    this.handleCancel = this.handleCancel.bind(this);
  }

  handleNameChange(event) {
    this.setState({ name: event.target.value });
  }

  handleDiscriptionChange(event) {
    this.setState({ discription: event.target.value });
  }

  handleOK(event) {
    event.preventDefault();
    this.props.dispatch(ServerGroupAction.createServerGroup({
      Name: this.state.name,
      Discription: this.state.discription
    }));
  }

  handleCancel(event) {
    event.preventDefault();
    this.props.dispatch(ServerGroupAction.closeCreateServerGroupDialog());
  }

  // Why do we pass action here?
  // Because I don't know how to get form content here if pass a function.
  render() {
    if (this.props.serverApp.openCreateServerGroupDialog) {
      return (
        <DialogFrame>
          <DialogTitle value="Create Server Group" />
          <DialogHR />
          <DialogContentDiv>
            <p>Name</p>
            <input onChange={this.handleNameChange}/>
            <p>Discription</p>
            <input onChange={this.handleDiscriptionChange}/>
          </DialogContentDiv>
          <DialogHR />
          <DialogControlDiv>
            <DialogButton name="Cancel" onClick={this.handleCancel} />
            <DialogButton name="OK" onClick={this.handleOK} />
          </DialogControlDiv>
        </DialogFrame>
      );
    }
    return (<div />);
  }
}

function mapStateToProps(state) {
  const { serverApp } = state;
  return { serverApp };
}

CreateServerGroupDialog.propTypes = {
  dispatch: PropTypes.func,
  serverApp: PropTypes.object,
  onOK: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(CreateServerGroupDialog, styles));
