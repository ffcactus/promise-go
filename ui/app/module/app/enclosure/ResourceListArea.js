import React from 'react';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import ResourceListHeadArea from './ResourceListHeadArea';
import ResourceListControlArea from './ResourceListControlArea';
import ResourceListContainer from './ResourceListContainer';
import styles from './App.css';

class ResourceListArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-column-container list-area border-row">
        <ResourceListHeadArea />
        <ResourceListControlArea />
        <ResourceListContainer />
      </div>
    );
  }
}

ResourceListArea.propTypes = {
};

export default connect()(CSSModules(ResourceListArea, styles, {allowMultiple: true}));

