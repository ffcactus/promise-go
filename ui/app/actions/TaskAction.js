import * as types from './types';
import * as Rest from '../utils/Rest';

export function getTaskListStart() {
  return {
    type: types.TASK_LIST_GET_BEGIN
  };
}

export function getTaskListSuccess(taskList) {
  return {
    type: types.TASK_LIST_GET_SUCCESS,
    info: taskList
  };
}

export function getTaskListFailure() {
  return {
    type: types.TASK_LIST_GET_FAILURE,
  };
}

export function getTaskList(hostname) {
  return dispatch => {
    dispatch(getTaskListStart());
    Rest.getTaskList(hostname).then((response) => {
      if (response.status === 200) {
        dispatch(getTaskListSuccess(response.response));
        return;
      }
      dispatch(getTaskListFailure());
    });
  };
}

function taskCreate(task) {
  return {
    type: types.TASK_CREATE,
    info: task
  };
}

function taskUpdate(task) {
  return {
    type: types.TASK_UPDATE,
    info: task
  };
}

function taskDelete(uri) {
  return {
    type: types.TASK_DELETE,
    info: uri
  };
}

export function taskMessage(event) {
  const task = JSON.parse(event.Message);
  switch (event.Type) {
    case 'Create':
      return taskCreate(task);
    case 'Update':
      return taskUpdate(task);
    case 'Delete':
      return taskDelete(event.Uri);
    default:
      // Should not be here.
      return taskUpdate(event.message);
  }
}

function message(eventString) {
  const event = JSON.parse(eventString);
  switch(event.Category) {
    case 'Task':
      return taskMessage(event);
    default:
      return null;
  }
}

function initTaskServiceStart() {
  return {
    type: types.INIT_TASK_SVC_START
  };
}

function initTaskServiceSuccess() {
  return {
    type: types.INIT_TASK_SVC_SUCCESS
  };
}

function initTaskServiceFailure() {
  return {
    type: types.INIT_TASK_SVC_FAILURE
  };
}

function taskServiceClose() {
  return {
    type: types.TASK_SVC_CLOSE
  };
}

// The action used when application start.
export function initTaskService(hostname) {
  return dispatch => {
    dispatch(initTaskServiceStart());
    // First, get task list.
    Rest.getTaskList(hostname, 0, -1).then((response) => {
      if (response.status === 200) {
        // dispatch(getTaskListSuccess(response.response));
        // Then subscribe message.
        const socket = new WebSocket('ws://' + hostname + ':8081' + '/director/rich/v1/task/ws');
        socket.onopen = () => {
          dispatch(initTaskServiceSuccess());
        };
        socket.onclose = () => {
          dispatch(taskServiceClose());
        };
        socket.onmessage = event => {
          dispatch(message(event.data));
        };
        socket.onerror = () => {
          dispatch(initTaskServiceFailure());
        };
      } else {
        dispatch(initTaskServiceFailure());
      }
    }).catch(() => {
      dispatch(initTaskServiceFailure());
    });
  };
}

