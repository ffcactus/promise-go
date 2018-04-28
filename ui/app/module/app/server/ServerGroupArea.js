import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import { ActionType } from './ConstValue';
import ServerGroupList from './ServerGroupList';
import ServerGroupControlArea from './ServerGroupControlArea';

function ServerGroupArea() {
  return (
    <div styleName="ServerGroupArea">
      <AppCollectionMenu action={{type: ActionType.APP_EXIT}} />
      <ServerGroupControlArea />
      <ServerGroupList />
    </div>
  );
}

export default CSSModules(ServerGroupArea, styles);

