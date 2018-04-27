import { ActionType } from './ConstValue';
import { createGetAction } from '../../../client/common';

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

function onServerDelete(server) {
  return {
    type: ActionType.ON_SERVER_DELETE,
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
 *
 * @param {string} uri The URI to get the server.
 */
export function restGet(uri) {
  return createGetAction(
    uri,
    ActionType.SERVER_REST_GET_START,
    ActionType.SERVER_REST_GET_SUCCESS,
    ActionType.SERVER_REST_GET_MESSAGE,
    ActionType.SERVER_REST_GET_EXCEPTION,
  );
}
