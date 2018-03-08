import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';
import AppCollectionMenu from './AppCollectionMenu';

function AppToolBar(props) {
  return (
    <div styleName="AppToolBar">
      <AppCollectionMenu />
      {props.children}
    </div>
  );
}

AppToolBar.propTypes = {
  children: PropTypes.object
};

export default CSSModules(AppToolBar, Styles);
