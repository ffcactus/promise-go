import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import ServerGroupList from './ServerGroupList';
import ServerGroupControlArea from './ServerGroupControlArea';

function ServerSearchArea() {
  return (
    <div styleName="ServerSearchArea">
      <input />
    </div>
  );
}

export default CSSModules(ServerSearchArea, styles);

