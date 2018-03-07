import React from 'react';
import AppFrame from '../../promise/app/AppFrame';
import CommonTwoColumn from '../../promise/app/CommonTwoColumn';
import ServerList from './ServerList';
function Server() {
  return (
    <AppFrame>
      <CommonTwoColumn left={<ServerList/>} />
    </AppFrame>
  );
}

export default Server;

