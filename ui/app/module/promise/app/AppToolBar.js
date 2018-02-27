import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';

function AppToolBar(props) {
  return (
    <div styleName="AppToolBar">
      {props.children}
    </div>
  );
}

AppToolBar.propTypes = {
  children: PropTypes.object
};

export default CSSModules(AppToolBar, Styles);
