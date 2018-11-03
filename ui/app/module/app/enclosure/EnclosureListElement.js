import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import EnclosureTask from './EnclosureTask';
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
    const { enclosure } = this.props;
    const currentStyle = 'flex-item bottom-border center-left-container selectable ' + (this.props.selected ? 'selected' : 'not-selected');
    let percentage = 100;
    if (enclosure.UI !== null && enclosure.UI.task !== null) {
      percentage = enclosure.UI.task.Percentage;
    }
    return (
      <div styleName={currentStyle} onClick={this.onSelect} style={{height: '39px'}}>
        <Health health={enclosure.Health}/>
        <div styleName="center-container">
          <p>{enclosure.Name}</p>
        </div>
        {percentage === 100 ? null : <EnclosureTask percentage={percentage}/>}
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
