import React from 'react';
import CSSModules from 'react-css-modules';
import ResourceSearch from './ResourceSearch';
import styles from './App.css';

class ResourceListHeadArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container bottom-border" style={{maxHeight: '40px'}}>
        <ResourceSearch/>
      </div>
    );
  }
}

export default CSSModules(ResourceListHeadArea, styles, {allowMultiple: true});
