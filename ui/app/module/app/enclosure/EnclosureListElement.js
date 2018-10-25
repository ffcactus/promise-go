import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import * as EnclosureAction from './EnclosureAction';
import { Health } from '../../promise/common/Widget/Health';
import styles from './App.css';

class EnclosureListElement extends React.Component {
  constructor(props) {
    super(props);
    this.onSelect = this.onSelect.bind(this);
  }

  onSelect(event) {
    event.preventDefault();
    this.props.dispatch(EnclosureAction.selectElement(this.props.enclosure.URI));
  }

  render() {
    const currentStyle = 'flex-item button-border center-left-container selectable ' + (this.props.selected ? 'selected' : 'not-selected');
    return (
      <div styleName={currentStyle} onClick={this.onSelect} style={{height: '39px'}}>
        <Health health={this.props.enclosure.Health}/>
        <div styleName="center-container">
          <p>{this.props.enclosure.Name}</p>
        </div>
      </div>
    );
  }
}

EnclosureListElement.propTypes = {
  enclosure: PropTypes.object,
  selected: PropTypes.bool,
  dispatch: PropTypes.func,
};

function mapStateToProps(state, ownProps) {
  return {
    selected: ownProps.enclosure.URI === state.enclosureApp.enclosureUri,
  };
}

export default connect(mapStateToProps)(CSSModules(EnclosureListElement, styles, {allowMultiple: true}));
