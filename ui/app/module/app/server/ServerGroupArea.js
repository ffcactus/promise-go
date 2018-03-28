import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import ServerGroupList from './ServerGroupList';
import ServerGroupControlArea from './ServerGroupControlArea';

function ServerGroupArea() {
  return (
    <div styleName="ServerGroupArea">
      <AppCollectionMenu />
      <ServerGroupControlArea />
      <ServerGroupList />
    </div>
  );
}

export default CSSModules(ServerGroupArea, styles);

