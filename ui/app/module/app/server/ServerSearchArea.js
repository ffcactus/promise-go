import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

function ServerSearchArea() {
  return (
    <div styleName="ServerSearchArea">
      <input />
    </div>
  );
}

export default CSSModules(ServerSearchArea, styles);

