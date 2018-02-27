import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Startup from '../components/Startup';


class StartupContainer extends Component {

  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Startup startup={this.props.startup} />
    );
  }
}

function mapStateToProps(state) {
  const { startup } = state;
  return { startup };
}

StartupContainer.propTypes = {
  startup: PropTypes.object
};

export default connect(mapStateToProps)(StartupContainer);
