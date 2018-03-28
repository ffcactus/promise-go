import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';

function AppFrame(props) {
  return (
    <div styleName="AppFrame">
      {props.children}
    </div>
  );
}

AppFrame.propTypes = {
  children: PropTypes.object
};

export default CSSModules(AppFrame, Styles);
