import { EnclosureResource } from './ConstValue';
import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import * as EnclosureAction from './EnclosureAction';
import styles from './Enclosure.css';

class EnclosureResourceItem extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
    this.props.dispatch(EnclosureAction.select());
  }

  render() {
    const currentStyle = 'flex-item border-column selectable ' + (this.props.selected ? 'selected' : 'not-selected');
    return (
      <div styleName={currentStyle} onClick={this.onClick} style={{maxHeight: '40px'}}>
        <p>Enclosure</p>
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    selected: state.enclosureApp.currentResource === EnclosureResource.Enclosure
  };
}

EnclosureResourceItem.propTypes = {
  dispatch: PropTypes.func,
  selected: PropTypes.bool,
};


export default connect(mapStateToProps)(CSSModules(EnclosureResourceItem, styles, { allowMultiple: true }));
