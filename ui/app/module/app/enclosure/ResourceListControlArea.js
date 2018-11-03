import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './App.css';

class ResourceListControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container left-border" style={{maxHeight: '40px'}}/>
    );
  }
}

export default CSSModules(ResourceListControlArea, styles, {allowMultiple: true});
