import * as Client from './Client';
import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';

function appInitStart() {
  return {
    type: ActionType.APP_INIT_START,
  };
}

function appInitSuccess() {
  return {
    type: ActionType.APP_INIT_SUCCESS,
  };
}

function appInitFailure() {
  return {
    type: ActionType.APP_INIT_FAILURE,
  };
}

function getServerGroupListStart() {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_START
  };
}

function getServerGroupListSuccess(resp) {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_SUCCESS,
    info: resp,
  };
}

function getServerGroupListFailure() {
  return {
    type: ActionType.GET_SERVERGROUP_LIST_FAILURE,
  };
}

function getServerListStart() {
  return {
    type: ActionType.GET_SERVER_LIST_START
  };
}

function getServerListSuccess(resp) {
  return {
    type: ActionType.GET_SERVER_LIST_SUCCESS,
    info: resp
  };
}

function getServerListFailure() {
  return {
    type: ActionType.GET_SERVER_LIST_FAILURE
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

function onServerMessage(message) {
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

export function appInit(hostname) {
  WsAction.registerMessageAction('Server', onServerMessage);
  WsAction.registerMessageAction('ServerGroup', onServerGroupMessage);
  return (dispatch, state) => {
    dispatch(appInitStart());
    // First, we need get all the servergroup.
    dispatch(getServerGroupListStart());
    Client.getServerGroupList(hostname).then((sgResp) => {
      if (sgResp.status === 200) {
        dispatch(getServerGroupListSuccess(sgResp.response));
        // Then we get all the servers from the current servergroup.
        dispatch(getServerListStart());
        Client.getServerList(hostname, state.currentGroup).then((sResp) => {
          if (sResp.status === 200) {
            dispatch(getServerListSuccess(sResp.response));
            dispatch(appInitSuccess());
            return;
          }
          // If status code error in getting server list, init fails.
          dispatch(getServerListFailure());
          dispatch(appInitFailure());
        }).catch((e) => {
          // if exception raised in getting server list, init fails.
          dispatch(getServerListFailure(e));
          dispatch(appInitFailure(e));
        });
        return;
      }
      // if status code error in getting servergroup list, init fails.
      dispatch(getServerGroupListFailure());
      dispatch(appInitFailure());
    }).catch((e) => {
      // if exception raised in getting servergroup list, init fails.
      dispatch(getServerGroupListFailure(e));
      dispatch(appInitFailure(e));
    });
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
