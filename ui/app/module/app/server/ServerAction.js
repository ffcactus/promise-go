import { ActionType } from './ConstValue';
import { createGetAction } from '../../../client/common';

function onServerCreate(server) {
  return {
    type: ActionType.SERVER_WS_CREATE,
    info: server
  };
}

function onServerUpdate(server) {
  return {
    type: ActionType.SERVER_WS_UPDATE,
    info: server
  };
}

function onServerDelete(server) {
  return {
    type: ActionType.SERVER_WS_DELETE,
    info: server
  };
}

export function onServerMessage(message) {
  switch(message.Type) {
    case 'Create':
      return onServerCreate(message.Data);
    case 'Update':
      return onServerUpdate(message.Data);
    case 'Delete':
      return onServerDelete(message.Data);
    default:
      return {};
  }
}

export function uiListSelect(uri) {
  return {
    type: ActionType.SERVER_UI_LIST_SELECT,
    info: uri
  };
}

export function openAddServerDialog() {
  return {
    type: ActionType.SERVER_UI_DIALOG_ADD_OPEN
  };
}

export function closeAddServerDialog() {
  return {
    type: ActionType.SERVER_UI_DIALOG_ADD_CLOSE
  };
}

/**
 * When the server element did mount we get the server details.
 * @param {string} uri The URI to get the server.
 */
export function onElementDidMount(uri) {
  return createGetAction(
    uri,
    ActionType.SERVER_REST_GET_START,
    ActionType.SERVER_REST_GET_SUCCESS,
    ActionType.SERVER_REST_GET_MESSAGE,
    ActionType.SERVER_REST_GET_EXCEPTION,
  );
}

/**
 * When server list did mount, we need load all the server-servergroup depends on the selected servergroup.
 * @param {string} uri The URI to get the server.
 */
export function onListDidMount() {
  return (dispatch, getState) => {
    const uri = '/promise/v1/server-servergroup?$filter=ServerGroupID eq \'' + getState().serverApp.currentServerGroup.split('/').pop() + '\'';
    createGetAction(
      uri,
      ActionType.SSG_REST_GETLIST_START,
      ActionType.SSG_REST_GETLIST_SUCCESS,
      ActionType.SSG_REST_GETLIST_MESSAGE,
      ActionType.SSG_REST_GETLIST_EXCEPTION
    )(dispatch, getState);
  };
}

