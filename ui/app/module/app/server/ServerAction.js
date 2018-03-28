import * as Client from './Client';
import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';

export function getServerListStart() {
  return {
    type: ActionType.GET_SERVER_LIST_START
  };
}

export function getServerListSuccess(resp) {
  return {
    type: ActionType.GET_SERVER_LIST_SUCCESS,
    info: resp
  };
}

export function getServerListFailure() {
  return {
    type: ActionType.GET_SERVER_LIST_FAILURE
  };
}

function onServerCreate(server) {
  return {
    type: ActionType.ON_SERVER_CREATE,
    info: server
  };
}

function onServerUpdate(server) {
  return {
    type: ActionType.ON_SERVER_UPDATE,
    info: server
  };
}

function onServerDelete(id) {
  return {
    type: ActionType.ON_SERVER_DELETE,
    info: id
  };
}

export function onServerMessage(message) {
  switch(message.Type) {
    case 'Create':
      return onServerCreate(message.Data);
    case 'Update':
      return onServerUpdate(message.Data);
    case 'Delete':
      return onServerDelete(message.ResourceID);
    default:
      return {};
  }
}

function onServerGroupMessage(message) {
  switch(message.Type) {
    case 'Create':
      return onServerGroupCreate(message.Data);
    case 'Update':
      return onServerGroupUpdate(message.Data);
    case 'Delete':
      return onServerGroupDelete(message.ResourceID);
    case 'DeleteCollection':
      return onServerGroupDeleteCollection();
    default:
      return {};
  }
}

export function openAddServerDialog() {
  return {
    type: ActionType.OPEN_ADD_SERVER_DIALOG
  };
}

export function closeAddServerDialog() {
  return {
    type: ActionType.CLOSE_ADD_SERVER_DIALOG
  };
}

function getServerStart() {
  return {
    type: ActionType.GET_SERVER_START
  };
}

function getServerSuccess(resp) {
  return {
    type: ActionType.GET_SERVER_SUCCESS,
    info: resp
  };
}

function getServerFailure() {
  return {
    type: ActionType.GET_SERVER_FAILURE
  };
}

export function getServer(hostname, uri) {
  return (dispatch) => {
    dispatch(getServerStart());
    Client.getServer(hostname, uri).then((resp) => {
      if (resp.status === 200) {
        dispatch(getServerSuccess(resp.response));
        return;
      }
      dispatch(getServerFailure());
    }).catch((e) => {
      dispatch(getServerFailure(e));
    });
  };
}
