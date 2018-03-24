import React from 'react';
import AppFrame from '../../promise/app/AppFrame';
import GroupCollectionApp from '../../promise/app/GroupCollectionApp/GroupCollectionApp';
import ServerList from './ServerList';
import ServerGroupArea from './ServerGroupArea';
import ServerDetail from './ServerDetail';

function Server() {
  return (
    <AppFrame>
      <GroupCollectionApp group={<ServerGroupArea/>} elementList={<ServerList/>} detail={<ServerDetail/>} />
    </AppFrame>
  );
}

export default Server;

