import React from 'react';
import CSSModules from 'react-css-modules';
import ResourceControlAdd from './ResourceControlAdd';
import ResourceControlEdit from './ResourceControlEdit';
import styles from './App.css';

class ResourceDetailControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container bottom-border" style={{maxHeight: '40px'}}>
        <ResourceControlAdd />
        <ResourceControlEdit />
      </div>
    );
  }
}

export default CSSModules(ResourceDetailControlArea, styles, {allowMultiple: true});
