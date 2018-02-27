import { DesktopAction } from './ConstValue';
import * as client from './Client';

function setInitStart() {
  return {
    type: DesktopAction.SET_INIT_START,
    info: {
      hostname: window.location.hostname
    }
  };
}

function setInitFailure() {
  return {
    type: DesktopAction.SET_INIT_FAILURE
  };
}

/**
 * Initialize the desktop.
 */
export function init() {
  return dispatch => {
    dispatch(setInitStart());
    client.getGlobal(window.location.hostname).then((resp) => {
      if (resp.status === 200) {
        return;
      }
      dispatch(setInitFailure());
    }).catch(() => {
      dispatch(setInitFailure());
    });
  };
}

/**
 * Set the desktop to show app collection.
 */
export function setAppCollection() {
  return {
    type: DesktopAction.SET_APP_COLLECTION
  };
}
