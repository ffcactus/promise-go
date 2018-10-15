import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { push } from 'react-router-redux';
import CSSModules from 'react-css-modules';
import DialogFrame from './dialog/DialogFrame';
import DialogTitle from './dialog/DialogTitle';
import DialogHR from './dialog/DialogHR';
import DialogContentDiv from './dialog/DialogContentDiv';
import DialogControlDiv from './dialog/DialogControlDiv';
import DialogButton from './dialog/DialogButton';
import styles from './AppInitFailure.css';

class AppInitFailure extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
    this.props.dispatch(push('/'));
  }

  render() {
    return (
      <DialogFrame>
        <DialogTitle value="App Initialization Failure" />
        <DialogHR />
        <DialogContentDiv>
          <p>App Initialization Failure. Make sure the service is avaliable and try again.</p>
        </DialogContentDiv>
        <DialogHR />
        <DialogControlDiv>
          <DialogButton name="OK" onClick={this.onClick} />
        </DialogControlDiv>
      </DialogFrame>
    );
  }
}

AppInitFailure.propTypes = {
  dispatch: PropTypes.func,
};

export default connect()(CSSModules(AppInitFailure, styles));
