import { push } from 'react-router-redux';
import * as Client from './Client';
import { ActionType } from './ConstValue';
import * as DesktopAction from '../desktop/DesktopAction';
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
function loginRequest(hostname, username, password) {
  return {
    type: ActionType.LOGIN_START,
    hostname,
    username,
    password
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
function loginSuccess(token) {
  return {
    type: ActionType.LOGIN_SUCCESS,
    token
  };
}

/**
 * The async action of login. It will involve sync actions.
 * @param {string} username
 * @param {string} password
 * @param {string} afterLoginPath
 */
function login(hostname, username, password) {
  return (dispatch, state) => {
    dispatch(loginRequest(hostname, username, password));
    Client.login(hostname, username, password).then((response) => {
      if (response.status === 200) {
        dispatch(loginSuccess(response.response.token));
        // TODO
        // Is it good to do redirection in action?
        // browserHistory.push(afterLoginPath);
        console.info('from = ' + state.from);
        dispatch(push('/'));
        dispatch(DesktopAction.setAppCollection());
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
