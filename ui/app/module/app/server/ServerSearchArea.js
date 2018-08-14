import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

function ServerSearchArea() {
  return (
    <div styleName="ColumnFlexItem PromiseBoarder" style={{height: '40px'}}>
      <input />
    </div>
  );
}

export default CSSModules(ServerSearchArea, styles, {allowMultiple: true});

