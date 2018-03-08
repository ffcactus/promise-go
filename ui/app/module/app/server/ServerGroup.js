import React from 'react';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';
import ServerGroupDetail from './ServerGroupDetail';

function ServerGroup() {
  return (
    <React.Fragment>
      <AppCollectionMenu />
      <ServerGroupDetail />
    </React.Fragment>
  );
}

export default ServerGroup;

