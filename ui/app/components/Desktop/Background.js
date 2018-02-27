import React from 'react';
import CSSModules from 'react-css-modules';
import styles from '../../styles/Desktop/Desktop.css';


function Background() {
  return (
    <div styleName="background" />
  );
}

export default CSSModules(Background, styles);
