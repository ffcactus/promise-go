import * as Client from './Client';
import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';

export function getServerGroupListStart() {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_START
  };
}

export function getServerGroupListSuccess(resp) {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_SUCCESS,
    info: resp,
  };
}

export function getServerGroupListFailure() {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_FAILURE,
  };
}

function onServerGroupCreate(server) {
  return {
    type: ActionType.ON_SERVERGROUP_CREATE,
    info: server
  };
}

function onServerGroupUpdate(server) {
  return {
    type: ActionType.ON_SERVERGROUP_UPDATE,
    info: server
  };
}

function onServerGroupDelete(id) {
  return {
    type: ActionType.ON_SERVERGROUP_DELETE,
    info: id
  };
}

function onServerGroupDeleteCollection() {
  return {
    type: ActionType.ON_SERVERGROUP_DELETE_COLLECTION
  };
}

export function onServerGroupMessage(message) {
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
    const state = getState();
    const hostname = state.session.hostname;
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

export function onServerGroupSelected(name) {
  return {
    type: ActionType.ON_SERVERGROUP_SELECTED,
    info: name
  };
}
