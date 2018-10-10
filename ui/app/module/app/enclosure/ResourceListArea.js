import React from 'react';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import ResourceListHeadArea from './ResourceListHeadArea';
import ResourceListControlArea from './ResourceListControlArea';
import ResourceList from './ResourceList';
import styles from './Enclosure.css';

class ResourceListArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-column-container list-area border-row">
        <ResourceListHeadArea />
        <ResourceListControlArea />
        <ResourceList />
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
  };
}

ResourceListArea.propTypes = {
};

export default connect(mapStateToProps)(CSSModules(ResourceListArea, styles, {allowMultiple: true}));

