import React from 'react';
import GroupFrame from '../../promise/common/frame/group/GroupFrame';
import AppListArea from './AppListArea';
import AppGroupArea from './AppGroupArea';
import AppDetailArea from './AppDetailArea';

function Enclosure() {
  return (
    <GroupFrame
      groupSection={<AppGroupArea />}
      listSection={<AppListArea />}
      detailSection={<AppDetailArea />} />
  );
}

export default Enclosure;
