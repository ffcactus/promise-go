import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import * as Action from './ServerAction';
import Form from '../../promise/common/Form';

class CreateServerGroupDialog extends React.Component {
  constructor(props) {
    super(props);
  }

  // Why do we pass action here?
  // Because I don't know how to get form content here if pass a function.
  render() {
    return (
      <Form
        title="Create Server Group"
        onOKAction={Action.createServerGroup}
        onCancelAction={Action.closeCreateServerGroupDialog} >
        <p>Name</p>
        <input />
        <p>Description</p>
        <input />
      </Form>
    );
  }
}

function mapStateToProps(state) {
  const { server } = state;
  return { server };
}

CreateServerGroupDialog.propTypes = {
  dispatch: PropTypes.func,
  server: PropTypes.object,
  onOK: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(CreateServerGroupDialog, styles));
