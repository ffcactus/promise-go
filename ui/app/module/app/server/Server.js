import React from 'react';
import AppFrame from '../../promise/app/AppFrame';
import CommonTwoColumn from '../../promise/app/CommonTwoColumn';
import ServerList from './ServerList';
import AppCollectionMenu from '../../promise/app/AppCollectionMenu';

function Server() {
  return (
    <AppFrame>
      <CommonTwoColumn left={
        <React.Fragment>
          <AppCollectionMenu />
          <ServerList/>
        </React.Fragment>} />
    </AppFrame>
  );
}

export default Server;

