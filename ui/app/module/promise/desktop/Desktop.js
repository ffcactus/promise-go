import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Desktop.css';


function Desktop(props) {
  return (
    <div styleName="desktop">
      {props.children}
    </div>
  );
}

Desktop.propTypes = {
  children: PropTypes.array
};

export default CSSModules(Desktop, styles);
