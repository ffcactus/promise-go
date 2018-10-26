import { ActionType } from './ConstValue';

function onCreate(task) {
  return {
    type: ActionType.TASK_WS_CREATE,
    info: task
  };
}

function onUpdate(task) {
  return {
    type: ActionType.TASK_WS_UPDATE,
    info: task
  };
}

function onDelete(task) {
  return {
    type: ActionType.TASK_WS_DELETE,
    info: task
  };
}

export function onTaskMessage(message) {
  switch (message.Type) {
    case 'Create':
      return onCreate(message.Data);
    case 'Update':
      return onUpdate(message.Data);
    case 'Delete':
      return onDelete(message.Data);
    default:
      return {};
  }
}
