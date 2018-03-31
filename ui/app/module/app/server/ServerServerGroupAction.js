import { ActionType } from './ConstValue';

function onServerServerGroupCreate(ssg) {
  return {
    type: ActionType.ON_SERVER_SERVERGROUP_CREATE,
    info: ssg
  };
}

function onServerServerGroupUpdate(ssg) {
  return {
    type: ActionType.ON_SERVER_SERVERGROUP_UPDATE,
    info: ssg
  };
}

function onServerServerGroupDelete(id) {
  return {
    type: ActionType.ON_SERVER_SERVERGROUP_DELETE,
    info: id
  };
}

export function onServerServerGroupMessage(message) {
  switch(message.Type) {
    case 'Create':
      return onServerServerGroupCreate(message.Data);
    case 'Update':
      return onServerServerGroupUpdate(message.Data);
    case 'Delete':
      return onServerServerGroupDelete(message.ResourceID);
    default:
      return {};
  }
}
