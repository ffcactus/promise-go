import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import * as ServerAction from './ServerAction';
import DialogFrame from '../../promise/common/dialog/DialogFrame';
import DialogTitle from '../../promise/common/dialog/DialogTitle';
import DialogHR from '../../promise/common/dialog/DialogHR';
import DialogContentDiv from '../../promise/common/dialog/DialogContentDiv';
import DialogControlDiv from '../../promise/common/dialog/DialogControlDiv';
import DialogButton from '../../promise/common/dialog/DialogButton';


class DiscoverServerDialog extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      discription: '',
      hostname: '',
      username: '',
      password: ''
    };
    this.onNameChange = this.onNameChange.bind(this);
    this.onDiscriptionChange = this.onDiscriptionChange.bind(this);
    this.onHostnameChange = this.onHostnameChange.bind(this);
    this.onUsernameChange = this.onUsernameChange.bind(this);
    this.onPasswordChange = this.onPasswordChange.bind(this);
    this.handleOK = this.handleOK.bind(this);
    this.handleCancel = this.handleCancel.bind(this);
  }

  onNameChange(event) {
    this.setState({ name: event.target.value });
  }

  onDiscriptionChange(event) {
    this.setState({ discription: event.target.value });
  }

  onHostnameChange(event) {
    this.setState({ hostname: event.target.value });
  }

  onUsernameChange(event) {
    this.setState({ username: event.target.value });
  }

  onPasswordChange(event) {
    this.setState({ password: event.target.value });
  }

  handleOK(event) {
    event.preventDefault();
    this.props.dispatch(ServerAction.discoverServer({
      Name: this.state.name,
      Discription: this.state.discription,
      Hostname: this.state.hostname,
      Username: this.state.username,
      Password: this.state.password
    }));
  }

  handleCancel(event) {
    event.preventDefault();
    this.props.dispatch(ServerAction.closeDiscoverServerDialog());
  }

  // Why do we pass action here?
  // Because I don't know how to get form content here if pass a function.
  render() {
    if (this.props.openDiscoverServerDialog) {
      return (
        <DialogFrame>
          <DialogTitle value="Discover Server" />
          <DialogHR />
          <DialogContentDiv>
            <p>Name</p>
            <input onChange={this.onNameChange}/>
            <p>Discription</p>
            <input onChange={this.onDiscriptionChange}/>
            <p>Hostname</p>
            <input onChange={this.onHostnameChange}/>
            <p>Username</p>
            <input onChange={this.onUsernameChange}/>
            <p>Password</p>
            <input onChange={this.onPasswordChange}/>
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
  return { openDiscoverServerDialog: state.serverApp.openDiscoverServerDialog };
}

DiscoverServerDialog.propTypes = {
  dispatch: PropTypes.func,
  openDiscoverServerDialog: PropTypes.bool,
  onOK: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(DiscoverServerDialog, styles));
