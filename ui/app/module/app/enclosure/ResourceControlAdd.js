import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import DiscoverEnclosureDialog from './DiscoverEnclosureDialog';
import TestDialog from './TestDialog';
import * as EnclosureAction from './EnclosureAction';
import CenterDiv from '../../promise/common/CenterDiv';
import styles from './App.css';

/**
 * ResourceControlAdd is the button to pop a dialog for server adding.
 */
class ResourceControlAdd extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
    if (! this.props.openDiscoverEnclosureDialog) {
      this.props.dispatch(EnclosureAction.openDiscoverDialog());
    }
  }

  render() {
    const icon = require('../../promise/common/img/icon/Navigation_Add_2x.png');
    return (
      <div styleName="main-control-button-container" style={{float: 'right'}}>
        <TestDialog />
        <CenterDiv><img src={icon} onClick={this.onClick} style={{display: 'block', margin: 'auto', height: '30px'}}/></CenterDiv>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    openDiscoverEnclosureDialog: state.enclosureApp.openDiscoverEnclosureDialog
  };
}

ResourceControlAdd.propTypes = {
  dispatch: PropTypes.func,
  openDiscoverEnclosureDialog: PropTypes.bool,
};

export default connect(mapStateToProps)(CSSModules(ResourceControlAdd, styles));
