import React from 'react';
import AppFrame from '../../promise/app/AppFrame';
import GroupCollectionApp from '../../promise/app/GroupCollectionApp/GroupCollectionApp';
import ServerListArea from './ServerListArea';
import ServerGroupArea from './ServerGroupArea';
import ServerDetailContainer from './ServerDetailContainer';

function Server() {
  return (
    <AppFrame>
      <GroupCollectionApp group={<ServerGroupArea/>} element={<ServerListArea/>} detail={<ServerDetailContainer/>} />
    </AppFrame>
  );
}

export default Server;

