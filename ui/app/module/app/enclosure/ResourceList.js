import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Enclosure.css';

class ResourceList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container border-column flex-item-last" />
    );
  }
}

export default CSSModules(ResourceList, styles, {allowMultiple: true});
