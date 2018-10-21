import React from 'react';
import CSSModules from 'react-css-modules';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import { ActionType } from './ConstValue';
import ServerGroupList from './ServerGroupList';
import ServerGroupControlArea from './ServerGroupControlArea';
import styles from './App.css';

function ServerGroupArea() {
  return (
    <div styleName="flex-column-container group-area">
      <AppCollectionMenu action={{type: ActionType.APP_SERVER_EXIT}} />
      <ServerGroupControlArea />
      <ServerGroupList />
    </div>
  );
}

export default CSSModules(ServerGroupArea, styles, {allowMultiple: true});

