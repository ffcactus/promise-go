import * as types from './types';
import * as Rest from '../utils/Rest';

export function getServerListStart() {
  return {
    type: types.SERVER_LIST_GET_BEGIN
  };
}

export function getServerListSuccess(list) {
  return {
    type: types.SERVER_LIST_GET_SUCCESS,
    info: list
  };
}

export function getServerListFailure() {
  return {
    type: types.SERVER_LIST_GET_FAILURE,
  };
}

function serverCreate(server) {
  return {
    type: types.SERVER_CREATE,
    info: server
  };
}

function serverUpdate(server) {
  return {
    type: types.SERVER_UPDATE,
    info: server
  };
}

function serverDelete(uri) {
  return {
    type: types.SERVER_DELETE,
    info: uri
  };
}


export function serverMessage(event) {
  const server = JSON.parse(event.Message);
  switch (event.Type) {
    case 'Create':
      return serverCreate(server);
    case 'Update':
      return serverUpdate(server);
    case 'Delete':
      return serverDelete(event.Uri);
    default:
      // Should not be here.
      return serverUpdate(event.message);
  }
}

function message(eventString) {
  const event = JSON.parse(eventString);
  switch(event.Category) {
    case 'Server':
      return serverMessage(event);
    default:
      return null;
  }
}

function initServerServiceStart() {
  return {
    type: types.INIT_SERVER_SVC_START
  };
}

function initServerServiceSuccess() {
  return {
    type: types.INIT_SERVER_SVC_SUCCESS
  };
}

function initServerServiceFailure() {
  return {
    type: types.INIT_SERVER_SVC_FAILURE
  };
}

function serverServiceClose() {
  return {
    type: types.SERVER_SVC_CLOSE
  };
}

// The action used when application start.
export function initServerService(hostname) {
  return dispatch => {
    dispatch(initServerServiceStart());
    // First, get task list.
    Rest.getServerList(hostname, 0, -1).then((response) => {
      if (response.status === 200) {
        dispatch(getServerListSuccess(response.response));
        // Then subscribe message.
        const socket = new WebSocket('ws://' + hostname + ':8080' + '/director/rich/v1/server/ws');
        socket.onopen = () => {
          dispatch(initServerServiceSuccess());
        };
        socket.onclose = () => {
          dispatch(serverServiceClose());
        };
        socket.onmessage = event => {
          dispatch(message(event.data));
        };
        socket.onerror = () => {
          dispatch(initServerServiceFailure());
        };
      } else {
        dispatch(initServerServiceFailure());
      }
    });
  };
}
