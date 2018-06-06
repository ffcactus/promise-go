import React from 'react';
import AppFrame from '../../promise/app/AppFrame';
import GroupCollectionApp from '../../promise/app/GroupCollectionApp/GroupCollectionApp';
import ElementArea from './ElementArea';
import GroupArea from './GroupArea';
import DetailContainer from './ServerDetailContainer';

function Server() {
  return (
    <AppFrame>
      <GroupCollectionApp group={<GroupArea/>} element={<ElementArea/>} detail={<DetailContainer/>} />
    </AppFrame>
  );
}

export default Server;
