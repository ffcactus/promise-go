import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './ServerProfile.css';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import { ActionType } from './ConstValue';
import ModelList from './ModelList';

function GroupArea() {
  return (
    <div styleName="ServerGroupArea">
      <AppCollectionMenu action={{type: ActionType.APP_EXIT}} />
      <ModelList />
    </div>
  );
}

export default CSSModules(GroupArea, styles);
