import React from 'react';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import ServerGroupList from './ServerGroupList';
import ServerGroupControlArea from './ServerGroupControlArea';

function ServerGroupArea() {
  return (
    <React.Fragment>
      <AppCollectionMenu />
      <ServerGroupControlArea />
      <ServerGroupList />
    </React.Fragment>
  );
}

export default ServerGroupArea;

