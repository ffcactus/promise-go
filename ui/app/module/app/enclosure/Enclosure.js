import React from 'react';
import GroupFrame from '../../promise/common/frame/group/GroupFrame';
import ResourceListArea from './ResourceListArea';
import EnclosureGroupArea from './EnclosureGroupArea';
import EnclosureDetailArea from './EnclosureDetailArea';

function Enclosure() {
  return (
    <GroupFrame
      groupSection={<EnclosureGroupArea />}
      listSection={<ResourceListArea />}
      detailSection={<EnclosureDetailArea />} />
  );
}

export default Enclosure;
