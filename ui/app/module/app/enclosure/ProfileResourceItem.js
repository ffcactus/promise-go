import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import { EnclosureResource } from './ConstValue';
import * as ProfileAction from './ProfileAction';
import styles from './App.css';

class ProfileResourceItem extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
    this.props.dispatch(ProfileAction.selectResource());
  }


  render() {
    const currentStyle = 'center-container flex-item border-column selectable ' + (this.props.selected ? 'selected' : 'not-selected');
    return (
      <div styleName={currentStyle} onClick={this.onClick} style={{maxHeight: '40px'}}>
        <p>Profile</p>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    selected: state.enclosureApp.currentResource === EnclosureResource.Profile
  };
}

ProfileResourceItem.propTypes = {
  dispatch: PropTypes.func,
  selected: PropTypes.bool,
};


export default connect(mapStateToProps)(CSSModules(ProfileResourceItem, styles, { allowMultiple: true }));
