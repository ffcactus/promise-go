import * as ActionConst from '../constValue/action/Desktop';
import * as client from '../client/global';

function initStart() {
  return {
    type: ActionConst.DESKTOP_INIT_START,
    info: {
      hostname: window.location.hostname
    }
  };
}

function initSuccess() {
  return {
    type: ActionConst.DESKTOP_INIT_SUCCESS
  };
}

function initFailure() {
  return {
    type: ActionConst.DESKTOP_INIT_FAILURE
  };
}

export function init() {
  return dispatch => {
    dispatch(initStart());
    client.getGlobal(window.location.hostname).then((resp) => {
      if (resp.status === 200) {
        dispatch(initSuccess());
        return;
      }
      dispatch(initFailure());
    }).catch(() => {
      dispatch(initFailure());
    });
  };
}
