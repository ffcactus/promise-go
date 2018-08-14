import { ActionType } from './ConstValue';

function onTaskCreate(task) {
  return {
    type: ActionType.TASK_WS_CREATE,
    info: task
  };
}

function onTaskUpdate(task) {
  return {
    type: ActionType.TASK_WS_UPDATE,
    info: task
  };
}

function onTaskDelete(task) {
  return {
    type: ActionType.TASK_WS_DELETE,
    info: task
  };
}

export function onTaskMessage(message) {
  console.info(message.Data.Percentage);

  switch (message.Type) {
    case 'Create':
      return onTaskCreate(message.Data);
    case 'Update':
      return onTaskUpdate(message.Data);
    case 'Delete':
      return onTaskDelete(message.Data);
    default:
      return {};
  }
}
