import React from 'react';
import AppFrame from '../../promise/app/AppFrame';
import GroupCollectionApp from '../../promise/app/GroupCollectionApp/GroupCollectionApp';
import ServerList from './ServerList';
import ServerGroup from './ServerGroup';
import ServerDetail from './ServerDetail';

function Server() {
  return (
    <AppFrame>
      <GroupCollectionApp group={<ServerGroup/>} elementList={<ServerList/>} detail={<ServerDetail/>} />
    </AppFrame>
  );
}

export default Server;

