import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Desktop from './Desktop';
import Background from './Background';
import AppCollection from './AppCollection';

class DesktopContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Desktop>
        <Background />
        <AppCollection />
      </Desktop>
    );
  }
}

function mapStateToProps(state) {
  const { desktop } = state;
  return { desktop };
}

DesktopContainer.propTypes = {
  desktop: PropTypes.object,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(DesktopContainer);
