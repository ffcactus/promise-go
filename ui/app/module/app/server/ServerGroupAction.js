import * as Client from './Client';
import { ActionType } from './ConstValue';

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

function onServerGroupCreate(sg) {
  return {
    type: ActionType.ON_SERVERGROUP_CREATE,
    info: sg
  };
}

function onServerGroupUpdate(sg) {
  return {
    type: ActionType.ON_SERVERGROUP_UPDATE,
    info: sg
  };
}

function onServerGroupDelete(sg) {
  return {
    type: ActionType.ON_SERVERGROUP_DELETE,
    info: sg
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
 * This action will be called when user clicks on a servergroup.
 * @param {string} id The ID of the group been selected.
 */
export function onServerGroupSelected(servergroup) {
  return (dispatch, getState) => {
    const hostname = getState().session.hostname;

    dispatch({
      type: ActionType.ON_SERVERGROUP_SELECTED,
      info: servergroup
    });
    dispatch({
      type: ActionType.GET_SERVER_LIST_START
    });
    // Get server list start.
    Client.getServerListByGroup(hostname, servergroup.ID).then((resp) => {
      if (resp.status === 200) {
        // Get server list success.
        dispatch({
          type: ActionType.GET_SERVER_LIST_SUCCESS,
          info: resp.response
        });
        return;
      }
      // Get server list failure.
      dispatch({
        type: ActionType.GET_SERVER_LIST_FAILURE,
        info: resp.response
      });
    }).catch(() => {
      // Get server list failure.
      createServerGroupFailure({
        type: ActionType.GET_SERVER_LIST_FAILURE,
      });
    });
  };
}
