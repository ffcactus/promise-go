import React from 'react';
import CSSModules from 'react-css-modules';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import EnclosureResourceItem from './EnclosureResourceItem';
import ProfileResourceItem from './ProfileResourceItem';
import IDPoolResourceItem from './IDPoolResourceItem';
import { ActionType } from './ConstValue';
import styles from './App.css';

function AppGroupArea() {
  return (
    <div styleName="flex-column-container group-area">
      <AppCollectionMenu action={{type: ActionType.APP_ENCLOSURE_EXIT}} />
      <EnclosureResourceItem />
      <ProfileResourceItem />
      <IDPoolResourceItem />
    </div>
  );
}

export default CSSModules(AppGroupArea, styles, {allowMultiple: true});