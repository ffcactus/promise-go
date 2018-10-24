import { ActionType } from './ConstValue';
import { createGetAction, createPostAction } from '../../promise/common/Client';

/**
 * This action will be called when user select enclosure resource..
 */
export function selectResource() {
  return {
    type: ActionType.ENCLOSURE_UI_SELECT_RESOURCE,
  };
}

/**
 * This action means user open discover enclosure dialog.
 */
export function openDiscoverDialog() {
  return {
    type: ActionType.ENCLOSURE_UI_DIALOG_DISCOVER_OPEN
  };
}

/**
 * This action means user closure discover enclosure dialog.
 */
export function closeDiscoverDialog() {
  return {
    type: ActionType.ENCLOSURE_UI_DIALOG_DISCOVER_CLOSE
  };
}

/**
 * This action means user select a single enclosure.
 * @param {string} uri The URI of the enclosure.
 */
export function selectElement(uri) {
  return (dispatch, getState) => {
    dispatch({
      type: ActionType.ENCLOSURE_UI_SELECT,
      info: uri
    });
    createGetAction(
      uri,
      ActionType.ENCLOSURE_REST_GET_START,
      ActionType.ENCLSOURE_REST_GET_SUCCESS,
      ActionType.ENCLSOURE_REST_GET_MESSAGE,
      ActionType.ENCLSOURE_REST_GET_EXCEPTION,
    )(dispatch, getState);
  };
}

/**
 * When user clicks on OK in discover enclosure dialog.
 * @param {object} request The request DTO.
 */
export function discover(request) {
  return (dispatch, getState) => {
    createPostAction(
      '/promise/v1/enclosure/action/discover',
      request,
      ActionType.ENCLOSURE_REST_DISCOVER_START,
      ActionType.ENCLOSURE_REST_DISCOVER_SUCCESS,
      ActionType.ENCLOSURE_REST_DISCOVER_MESSAGE,
      ActionType.ENCLOSURE_REST_DISCOVER_EXCEPTION,
    )(dispatch, getState);
  };
}

/**
 * This action is called on WS message
 * @param {object} message The DTO from WS.
 */
export function onMessage(message) {
  switch(message.Type) {
    case 'Create':
      return {
        type: ActionType.ENCLOSURE_WS_CREATE,
        info: message.Data
      };
    case 'Update':
      return {
        type: ActionType.ENCLOSURE_WS_UPDATE,
        info: message.Data
      };
    case 'Delete':
      return {
        type: ActionType.ENCLOSURE_WS_DELETE,
        info: message.Data
      };
    default:
      return {};
  }
}
