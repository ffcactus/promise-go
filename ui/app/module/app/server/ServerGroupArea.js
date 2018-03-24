import React from 'react';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import ServerGroupDetail from './ServerGroupDetail';
import ServerGroupControlArea from './ServerGroupControlArea';

function ServerGroupArea() {
  return (
    <React.Fragment>
      <AppCollectionMenu />
      <ServerGroupControlArea />
      <ServerGroupDetail />
    </React.Fragment>
  );
}

export default ServerGroupArea;

