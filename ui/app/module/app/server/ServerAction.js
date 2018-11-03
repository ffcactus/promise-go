import { ActionType } from './ConstValue';
import { createGetAction, createPostAction } from '../../promise/common/Client';

/**
 * This action is called on WS message
 * @param {object} message The DTO from WS.
 */
export function onServerMessage(message) {
  switch(message.Type) {
    case 'Create':
      return {
        type: ActionType.SERVER_WS_CREATE,
        info: message.Data
      };
    case 'Update':
      return {
        type: ActionType.SERVER_WS_UPDATE,
        info: message.Data
      };
    case 'Delete':
      return {
        type: ActionType.SERVER_WS_DELETE,
        info: message.Data
      };
    default:
      return {};
  }
}

export function openDiscoverServerDialog() {
  return {
    type: ActionType.SERVER_UI_DIALOG_DISCOVER_OPEN
  };
}

export function closeDiscoverServerDialog() {
  return {
    type: ActionType.SERVER_UI_DIALOG_DISCOVER_CLOSE
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
 * This action means user clicks on server in the list.
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

/**
 * When user clicks on OK in discover server dialog.
 * @param {object} request The request DTO.
 */
export function discoverServer(request) {
  return (dispatch, getState) => {
    createPostAction(
      '/promise/v1/server/action/discover',
      request,
      ActionType.SERVER_REST_DISCOVER_START,
      ActionType.SERVER_REST_DISCOVER_SUCCESS,
      ActionType.SERVER_REST_DISCOVER_MESSAGE,
      ActionType.SERVER_REST_DISCOVER_EXCEPTION,
    )(dispatch, getState);
  };
}
