import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';

function AppFrame(props) {
  return (
    <div styleName="AppFrame">
      <div styleName="Home">
        <Link to="/">Home</Link>
      </div>
      {props.children}
    </div>
  );
}

AppFrame.propTypes = {
  children: PropTypes.object
};

export default CSSModules(AppFrame, Styles);
