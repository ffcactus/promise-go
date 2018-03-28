import React from 'react';
import AppFrame from '../../promise/app/AppFrame';
import GroupCollectionApp from '../../promise/app/GroupCollectionApp/GroupCollectionApp';
import ServerListArea from './ServerListArea';
import ServerGroupArea from './ServerGroupArea';
import ServerDetailArea from './ServerDetailArea';

function Server() {
  return (
    <AppFrame>
      <GroupCollectionApp group={<ServerGroupArea/>} element={<ServerListArea/>} detail={<ServerDetailArea/>} />
    </AppFrame>
  );
}

export default Server;

