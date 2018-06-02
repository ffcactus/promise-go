import { ActionType } from './ConstValue';
import { createGetAction } from '../../promise/common/Client';

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
export function getServer(uri) {
  return createGetAction(
    uri,
    ActionType.SERVER_REST_GET_START,
    ActionType.SERVER_REST_GET_SUCCESS,
    ActionType.SERVER_REST_GET_MESSAGE,
    ActionType.SERVER_REST_GET_EXCEPTION,
  );
}

/**
 * When user clicks a differenct server, we need load the new server.
 * @param {string} uri The URI to get the server.
 */
export function uiListSelect(uri) {
  return (dispatch, getState) => {
    dispatch({
      type: ActionType.SERVER_UI_LIST_SELECT,
      info: uri
    });
    createGetAction(
      uri,
      ActionType.SERVER_REST_GET_START,
      ActionType.SERVER_REST_GET_SUCCESS,
      ActionType.SERVER_REST_GET_MESSAGE,
      ActionType.SERVER_REST_GET_EXCEPTION,
    )(dispatch, getState);
  };
}

/**
 * When user clicks orderBy selector.
 * @param {string} orderBy The by what to order the servers.
 */
export function onServerOrderChange(orderBy) {
  return {
    type: ActionType.SERVER_UI_ORDERBY_CHANGE,
    info: orderBy,
  };
}
