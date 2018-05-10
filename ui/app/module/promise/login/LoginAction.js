import { push } from 'react-router-redux';
import * as Client from './Client';
import { ActionType } from './ConstValue';
import * as DesktopAction from '../desktop/DesktopAction';
import * as WsAction from '../ws/WsAction';
const LOGIN_FAILURE_WAIT_TIME = 3000;

/*
 * When the user press login button, the following state takes in order:
 * 1. login start.
 * 2. wait login process.
 * 3. if login success, done.
 * 4. if login failure, wait a while to try again.
 */

/**
 * Login request start.
 * @param {string} username
 * @param {string} password
 */
function loginStart() {
  return {
    type: ActionType.LOGIN_START,
  };
}

/**
 * Login failed.
 * @param {object} info - The json response from server.
 */
function loginFailure(info) {
  return {
    type: ActionType.LOGIN_FAILURE,
    info
  };
}

/**
 * Login failed, and waited enough time to try again.
 */
function loginFailureTimeout() {
  return {
    type: ActionType.LOGIN_FAILURE_TIMEOUT
  };
}

/**
 * Login success.
 * @param {string} token
 */
function loginSuccess(hostname, username, token) {
  return {
    type: ActionType.LOGIN_SUCCESS,
    info: {
      hostname,
      username,
      token
    }
  };
}

/**
 * The async action of login. It will involve sync actions.
 * @param {string} username
 * @param {string} password
 * @param {string} afterLoginPath
 */
function login(hostname, username, password, from) {
  return (dispatch) => {
    dispatch(loginStart());
    Client.login(hostname, username, password).then((response) => {
      if (response.status === 200) {
        dispatch(WsAction.createWsConnection(hostname));
        dispatch(loginSuccess(hostname, username, response.response.token));
        // TODO
        // Is it good to do redirection in action?
        // browserHistory.push(afterLoginPath);
        dispatch(DesktopAction.setAppCollection());
        dispatch(push(from));
        return;
      }
      setTimeout(() => {
        dispatch(loginFailureTimeout());
      }, LOGIN_FAILURE_WAIT_TIME);
      dispatch(loginFailure(Array.isArray(response.response) ? response.response[0] : response.response));
    }).catch((reason) => {
      setTimeout(() => {
        dispatch(loginFailureTimeout());
      }, LOGIN_FAILURE_WAIT_TIME);
      dispatch(loginFailure(reason.message));
    });
  };
}

export { login };
