import * as types from './types';
import * as Rest from '../utils/Rest';

export function hardwareActionPopAddDialog() {
  return {
    type: types.HARDWARE_POP_ADD_DIALOG
  };
}

export function hardwareActionDialogCancel() {
  return {
    type: types.HARDWARE_ADD_DIALOG_CANCEL
  };
}

export function hardwareActionDialogOk(hardware) {
  return {
    type: types.HARDWARE_ADD_DIALOG_OK,
    info: hardware
  };
}

export function serverListLoadBegin() {
  return {
    type: types.SERVER_LIST_LOAD_BEGIN
  };
}

export function serverListLoadFailure() {
  return {
    type: types.SERVER_LIST_LOAD_FAILURE
  };
}

export function serverListLoadSuccess(serverCollection) {
  return {
    type: types.SERVER_LIST_LOAD_SUCCESS,
    info: serverCollection
  };
}

export function serverLoadBegin() {
  return {
    type: types.SERVER_GET_BEGIN,
  };
}

export function serverLoadSuccess(server) {
  return {
    type: types.SERVER_GET_SUCCESS,
    info: server
  };
}

export function serverLoadFailure() {
  return {
    type: types.SERVER_GET_FAILURE
  };
}

export function loadAllServer(hostname) {
  return dispatch => {
    dispatch(serverListLoadBegin());
    Rest.getServerList(hostname).then((response) => {
      if (response.status === 200) {
        dispatch(serverListLoadSuccess(response.response));
        return;
      }
      dispatch(serverListLoadFailure());
    });
  };
}

export function loadServer(id) {
  return (dispatch, getState) => {
    const {hostAddress} = getState().global;
    dispatch(serverLoadBegin());
    Rest.loadServer(hostAddress, id).then((response) => {
      if (response.status === 200) {
        dispatch(serverLoadSuccess(response.response));
        return;
      }
      dispatch(serverLoadFailure());
    });
  };
}

