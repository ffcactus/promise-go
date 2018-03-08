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
  return (dispatch) => {
    dispatch(appInitStart());
    dispatch(getServerListStart());
    Client.getServerList(hostname).then((resp) => {
      if (resp.status === 200) {
        dispatch(appInitSuccess());
        dispatch(getServerListSuccess(resp.response));
        return;
      }
      dispatch(appInitFailure());
      dispatch(getServerListFailure());
    }).catch((e) => {
      dispatch(getServerListFailure(e));
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
