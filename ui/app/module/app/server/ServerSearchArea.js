import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

function ServerSearchArea() {
  return (
    <div styleName="flex-item flex-row-container border-column-first" style={{maxHeight: '40px'}}>
      <input />
    </div>
  );
}

export default CSSModules(ServerSearchArea, styles, {allowMultiple: true});

