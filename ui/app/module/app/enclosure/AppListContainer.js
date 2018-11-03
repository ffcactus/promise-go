import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { EnclosureResource } from './ConstValue';
import EnclosureList from './EnclosureList';
import IDPoolList from './IDPoolList';
import ProfileList from './ProfileList';

class AppListContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    switch(this.props.currentResource) {
      case EnclosureResource.Enclosure:
        return <EnclosureList/>;
      case EnclosureResource.IDPool:
        return <IDPoolList/>;
      case EnclosureResource.Profile:
        return <ProfileList/>;
      default:
        return <p>Unknown Resource List</p>;
    }
  }
}

AppListContainer.propTypes = {
  currentResource: PropTypes.string
};

function mapStateToProps(state) {
  return {
    currentResource: state.enclosureApp.currentResource
  };
}

export default connect(mapStateToProps)(AppListContainer);
