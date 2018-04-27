import * as Client from './Client';
import { ActionType } from './ConstValue';
import { createGetAction } from '../../../client/common';

export function getCollectionStart() {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_START
  };
}

export function getCollectionSuccess(resp) {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_SUCCESS,
    info: resp,
  };
}

export function getCollectionFailure() {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_FAILURE,
  };
}

function onServerGroupCreate(sg) {
  return {
    type: ActionType.SG_WS_CREATE,
    info: sg
  };
}

function onServerGroupUpdate(sg) {
  return {
    type: ActionType.SG_WS_UPDATE,
    info: sg
  };
}

function onServerGroupDelete(sg) {
  return {
    type: ActionType.SG_WS_DELETE,
    info: sg
  };
}

function onServerGroupDeleteCollection() {
  return {
    type: ActionType.SG_WS_DELETE_LIST
  };
}

export function onServerGroupMessage(message) {
  switch(message.Type) {
    case 'Create':
      return onServerGroupCreate(message.Data);
    case 'Update':
      return onServerGroupUpdate(message.Data);
    case 'Delete':
      return onServerGroupDelete(message.Data);
    case 'DeleteCollection':
      return onServerGroupDeleteCollection();
    default:
      return {};
  }
}

export function openCreateServerGroupDialog() {
  return {
    type: ActionType.OPEN_CREATE_SERVERGROUP_DIALOG
  };
}

export function closeCreateServerGroupDialog() {
  return {
    type: ActionType.CLOSE_CREATE_SERVERGROUP_DIALOG
  };
}

function createServerGroupStart() {
  return {
    type: ActionType.CREATE_SERVERGROUP_START
  };
}

function createServerGroupSuccess(responseDto) {
  return {
    type: ActionType.CREATE_SERVERGROUP_SUCCESS,
    info: responseDto
  };
}

function createServerGroupFailure(messages) {
  return {
    type: ActionType.CREATE_SERVERGROUP_FAILURE,
    info: messages
  };
}

export function createServerGroup(servergroup) {
  return (dispatch, getState) => {
    const hostname = getState().session.hostname;
    dispatch(createServerGroupStart());
    Client.postServerGroup(hostname, servergroup).then((resp) => {
      if (resp.status === 201) {
        dispatch(createServerGroupSuccess(resp.response));
        dispatch(closeCreateServerGroupDialog());
        return;
      }
      dispatch(createServerGroupFailure(resp.response));
    }).catch(e => {
      createServerGroupFailure(e);
    });
  };
}

/**
 *
 * @param {int} top How much you would like to get.
 * @param {int} skip From where you would like to get.
 * @param {string} filter The filter you would like to use.
 */
export function getCollection() {
  return createGetAction(
    '/promise/v1/servergroup',
    ActionType.SG_REST_GETLIST_START,
    ActionType.SG_REST_GETLIST_SUCCESS,
    ActionType.SG_REST_GETLIST_MESSAGE,
    ActionType.SG_REST_GETLIST_EXCEPTION,
  );
}

/**
 * This action will be called when user selects a servergroup from list.
 * @param {string} uri The URI of the servergroup been selected.
 */
export function UiListSelect(uri) {
  return (dispatch, getState) => {
    dispatch({
      type: ActionType.SG_UI_LIST_SELECT,
      info: uri
    });
    // return an action and call it.
    createGetAction(
      '/promise/v1/server-servergroup?$filter=ServerGroupID eq \'' + uri.split('/').pop() + '\'',
      ActionType.SERVER_REST_GETLIST_START,
      ActionType.SERVER_REST_GETLIST_SUCCESS,
      ActionType.SERVER_REST_GETLIST_MESSAGE,
      ActionType.SERVER_REST_GETLIST_EXCEPTION,
    )(dispatch, getState);
  };
}
