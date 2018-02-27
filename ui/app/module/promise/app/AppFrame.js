import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';
import AppToolBar from './AppToolBar';
import AppMain from './AppMain';

function AppFrame(props) {
  return (
    <div styleName="AppFrame">
      <AppToolBar>
        <p>AppToolBar</p>
      </AppToolBar>
      <AppMain>
        {props.children}
      </AppMain>
    </div>
  );
}

AppFrame.propTypes = {
  children: PropTypes.object
};

export default CSSModules(AppFrame, Styles);
