import React from 'react';
import PropTypes from 'prop-types';
import ConfirmDialog from './common/ConfirmDialog';

class Example extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return <ConfirmDialog title="Initialization Failed" message="You need confirm the follow information."/>;
  }
}

Example.propTypes = {
  server: PropTypes.object,
  dispatch: PropTypes.func,
};

export default Example;
