import React from 'react';
import GroupFrame from '../../promise/common/frame/group/GroupFrame';
import EnclosureListArea from './EnclosureListArea';
import EnclosureGroupArea from './EnclosureGroupArea';
import EnclosureDetailArea from './EnclosureDetailArea';

function Enclosure() {
  return (
    <GroupFrame
      groupSection={<EnclosureGroupArea />}
      listSection={<EnclosureListArea />}
      detailSection={<EnclosureDetailArea />} />
  );
}

export default Enclosure;
