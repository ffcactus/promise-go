import React from 'react';
import CSSModules from 'react-css-modules';
import ServerDetailContainer from './ServerDetailContainer';
import ServerDetailControlArea from './ServerDetailControlArea';
import styles from './Server.css';

function ServerDetailArea() {
  return (
    <div styleName="flex-column-container detail-area border-row">
      <ServerDetailControlArea />
      <ServerDetailContainer />
    </div>
  );
}

export default CSSModules(ServerDetailArea, styles, {allowMultiple: true});

