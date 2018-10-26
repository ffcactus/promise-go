import React from 'react';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import ResourceListHeadArea from './ResourceListHeadArea';
import ResourceListControlArea from './ResourceListControlArea';
import AppListContainer from './AppListContainer';
import styles from './App.css';

class AppListArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-column-container list-area left-border">
        <ResourceListHeadArea />
        <ResourceListControlArea />
        <AppListContainer />
      </div>
    );
  }
}

AppListArea.propTypes = {
};

export default connect()(CSSModules(AppListArea, styles, {allowMultiple: true}));

