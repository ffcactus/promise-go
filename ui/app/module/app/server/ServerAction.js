import * as Client from './Client';
import { ActionType } from './ConstValue';

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
    type: ActionType.GET_SERVER_LIST_START
  };
}

function getServerSuccess(resp) {
  return {
    type: ActionType.GET_SERVER_LIST_SUCCESS,
    info: resp
  };
}

function getServerFailure() {
  return {
    type: ActionType.GET_SERVER_LIST_FAILURE
  };
}

export function appInit(hostname) {
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
