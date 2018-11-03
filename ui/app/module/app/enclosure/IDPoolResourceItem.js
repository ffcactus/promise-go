import { EnclosureResource } from './ConstValue';
import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import * as IDPoolAction from './IDPoolAction';
import styles from './App.css';

class IDPoolResourceItem extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
    this.props.dispatch(IDPoolAction.selectResource());
  }

  render() {
    const currentStyle = 'center-container flex-item bottom-border selectable ' + (this.props.selected ? 'selected' : 'not-selected');
    return (
      <div styleName={currentStyle} onClick={this.onClick} style={{maxHeight: '40px'}}>
        <p>ID Pool</p>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    selected: state.enclosureApp.currentResource === EnclosureResource.IDPool
  };
}

IDPoolResourceItem.propTypes = {
  dispatch: PropTypes.func,
  selected: PropTypes.bool,
};


export default connect(mapStateToProps)(CSSModules(IDPoolResourceItem, styles, { allowMultiple: true }));
