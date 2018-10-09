import React from 'react';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Enclosure.css';

class EnclosureListArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-column-container list-area border-row" />
    );
  }
}

function mapStateToProps(state) {
  return {
  };
}

EnclosureListArea.propTypes = {
};

export default connect(mapStateToProps)(CSSModules(EnclosureListArea, styles, {allowMultiple: true}));

