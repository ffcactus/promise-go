import React from 'react';
import GroupFrame from '../../promise/common/frame/group/GroupFrame';
import ServerListArea from './ServerListArea';
import ServerGroupArea from './ServerGroupArea';
import ServerDetailContainer from './ServerDetailContainer';

function Server() {
  return (
    <GroupFrame
      groupSection={<ServerGroupArea/>}
      listSection={<ServerListArea/>}
      detailSection={<ServerDetailContainer/>}/>
  );
}

export default Server;

