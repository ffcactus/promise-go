import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import * as EnclosureAction from './EnclosureAction';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogActions from '@material-ui/core/DialogActions';
import TextField from '@material-ui/core/TextField';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import styles from './App.css';

class TestDialog extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: '',
      description: '',
      address: '',
      username: '',
      password: '',
      type: 'E9000'
    };
    this.onNameChange = this.onNameChange.bind(this);
    this.onDiscriptionChange = this.onDiscriptionChange.bind(this);
    this.onAddressChange = this.onAddressChange.bind(this);
    this.onUsernameChange = this.onUsernameChange.bind(this);
    this.onPasswordChange = this.onPasswordChange.bind(this);
    this.onTypeChange = this.onTypeChange.bind(this);
    this.onOK = this.onOK.bind(this);
    this.onCancel = this.onCancel.bind(this);
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

  onOK(event) {
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

  onCancel(event) {
    event.preventDefault();
    this.props.dispatch(EnclosureAction.closeDiscoverDialog());
  }

  // Why do we pass action here?
  // Because I don't know how to get form content here if pass a function.
  render() {
    return (
      <Dialog
        aria-labelledby="Discover Enclosure"
        aria-describedby="Discover Enclosure"
        open={this.props.openDiscoverEnclosureDialog}
        onClose={this.onClose}
      >
        <DialogTitle id="Discover Enclosure">Discover Enclosure</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Input the login credential to the enclosure manager to add it to the system.
          </DialogContentText>
          <InputLabel htmlFor="Type">Type</InputLabel>
          <Select
            input={<Input name="Type" id="Type" fullWidth value={this.state.type} onChange={this.onTypeChange}/>}
          >
            <MenuItem value="E9000">E9000</MenuItem>
            <MenuItem value="Mock">Mock</MenuItem>
          </Select>
          <TextField autoFocus margin="dense" id="name" label="Name" fullWidth onChange={this.onNameChange}/>
          <TextField margin="dense" id="description" label="Description" fullWidth onChange={this.onDiscriptionChange}/>
          <TextField margin="dense" id="address" label="Address" fullWidth onChange={this.onAddressChange}/>
          <TextField margin="dense" id="username" label="Username" fullWidth onChange={this.onUsernameChange}/>
          <TextField margin="dense" id="username" label="Password" fullWidth type="password" onChange={this.onPasswordChange}/>
        </DialogContent>
        <DialogActions>
          <Button onClick={this.onCancel} color="primary">Cancel</Button>
          <Button onClick={this.onOK} color="primary">OK</Button>
        </DialogActions>
      </Dialog>
    );
  }
}

function mapStateToProps(state) {
  return { openDiscoverEnclosureDialog: state.enclosureApp.openDiscoverEnclosureDialog };
}

TestDialog.propTypes = {
  dispatch: PropTypes.func,
  styles: PropTypes.object,
  openDiscoverEnclosureDialog: PropTypes.bool,
  onOK: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(TestDialog, styles));
