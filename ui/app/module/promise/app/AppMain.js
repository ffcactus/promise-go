import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';

function AppMain(props) {
  return (
    <div styleName="AppMain">
      {props.children}
    </div>
  );
}

AppMain.propTypes = {
  children: PropTypes.object
};

export default CSSModules(AppMain, Styles);
