import React from 'react';
import GroupFrame from '../../promise/common/frame/group/GroupFrame';
import ServerListArea from './ServerListArea';
import ServerGroupArea from './ServerGroupArea';
import ServerDetailArea from './ServerDetailArea';

function Server() {
  return (
    <GroupFrame
      groupSection={<ServerGroupArea/>}
      listSection={<ServerListArea/>}
      detailSection={<ServerDetailArea/>}/>
  );
}

export default Server;

