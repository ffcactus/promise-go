import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import * as EnclosureAction from './EnclosureAction';
import DialogFrame from '../../promise/common/dialog/DialogFrame';
import DialogTitle from '../../promise/common/dialog/DialogTitle';
import DialogHR from '../../promise/common/dialog/DialogHR';
import DialogContentDiv from '../../promise/common/dialog/DialogContentDiv';
import DialogControlDiv from '../../promise/common/dialog/DialogControlDiv';
import DialogButton from '../../promise/common/dialog/DialogButton';
import Select from 'react-select';
import styles from './App.css';

const options = [
  {value: 'E9000', label: 'E9000'},
  {value: 'Mock', label: 'Mock'}
];

class DiscoverEnclosureDialog extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      description: '',
      address: '',
      username: '',
      password: ''
    };
    this.onNameChange = this.onNameChange.bind(this);
    this.onDiscriptionChange = this.onDiscriptionChange.bind(this);
    this.onAddressChange = this.onAddressChange.bind(this);
    this.onUsernameChange = this.onUsernameChange.bind(this);
    this.onPasswordChange = this.onPasswordChange.bind(this);
    this.onTypeChange = this.onTypeChange.bind(this);
    this.handleOK = this.handleOK.bind(this);
    this.handleCancel = this.handleCancel.bind(this);
  }

  onNameChange(event) {
    event.preventDefault();
    this.setState({ name: event.target.value });
  }

  onDiscriptionChange(event) {
    event.preventDefault();
    this.setState({ description: event.target.value });
  }

  onAddressChange(event) {
    event.preventDefault();
    this.setState({ address: event.target.value });
  }

  onUsernameChange(event) {
    event.preventDefault();
    this.setState({ username: event.target.value });
  }

  onPasswordChange(event) {
    event.preventDefault();
    this.setState({ password: event.target.value });
  }

  onTypeChange(event) {
    event.preventDefault();
    this.setState({ type: event.target.value });
  }

  handleOK(event) {
    event.preventDefault();
    this.props.dispatch(EnclosureAction.discover({
      Name: this.state.name,
      Description: this.state.description,
      Type: this.state.type,
      Address: this.state.address,
      Username: this.state.username,
      Password: this.state.password
    }));
  }

  handleCancel(event) {
    event.preventDefault();
    this.props.dispatch(EnclosureAction.closeDiscoverDialog());
  }

  // Why do we pass action here?
  // Because I don't know how to get form content here if pass a function.
  render() {
    if (this.props.openDiscoverEnclosureDialog) {
      return (
        <DialogFrame>
          <DialogTitle value="Discover Enclosure" />
          <DialogHR />
          <DialogContentDiv>
            <form>
              <label htmlFor="Name">Name</label>
              <input id="Name" type="text" aria-label="Name" aria-required="true" onChange={this.onNameChange}/>
              <br/>
              <label htmlFor="Discription">Discription</label>
              <input id="Discription" type="text" aria-label="Discription" onChange={this.onDiscriptionChange}/>
              <br/>
              <label htmlFor="Type">Type</label>
              <Select options={options} className="select"/>
              <br/>
              <label htmlFor="Address">Address</label>
              <input id="Address" type="text" aria-label="Address" aria-required="true" onChange={this.onAddressChange}/>
              <br/>
              <label htmlFor="Username">Username</label>
              <input id="Username" type="text" aria-label="Username" aria-required="true" onChange={this.onUsernameChange}/>
              <br/>
              <label htmlFor="Password">Password</label>
              <input id="Password" type="password" aria-label="Password" aria-required="true" onChange={this.onPasswordChange}/>
            </form>
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
  return { openDiscoverEnclosureDialog: state.enclosureApp.openDiscoverEnclosureDialog };
}

DiscoverEnclosureDialog.propTypes = {
  dispatch: PropTypes.func,
  openDiscoverEnclosureDialog: PropTypes.bool,
  onOK: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(DiscoverEnclosureDialog, styles));
