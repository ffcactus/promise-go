import * as types from './types';
import { browserHistory } from 'react-router';
import { startup } from './PlatformAction';
import { setHostAddress } from './GlobalAction';

// import * as Rest from '../utils/Rest';

// const LOGIN_FAILURE_WAIT_TIME = 3000;
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
    type: types.LOGIN_REQUEST,
    hostname,
    username,
    password
  };
}

/**
 * Login failed.
 * @param {object} info - The json response from server.
 */
// function loginFailure(info) {
//   return {
//     type: types.LOGIN_FAILURE,
//     info
//   };
// }

/**
 * Login failed, and waited enough time to try again.
 */
// function loginFailureTimeout() {
//   return {
//     type: types.LOGIN_FAILURE_TIMEOUT
//   };
// }

/**
 * Login success.
 * @param {string} token
 */
function loginSuccess(token) {
  return {
    type: types.LOGIN_SUCCESS,
    token
  };
}

/**
 * The async action of login. It will involve sync actions.
 * @param {string} username
 * @param {string} password
 * @param {string} afterLoginPath
 */
// function login(hostname, username, password, afterLoginPath) {
//   return dispatch => {
//     dispatch(loginRequest(hostname, username, password));
//     Rest.login(hostname, username, password).then((response) => {
//       if (response.status === 200) {
//         dispatch(loginSuccess(response.response.token));
//         // TODO
//         // Is it good to do redirection in action?
//         browserHistory.push(afterLoginPath);
//         return;
//       }
//       setTimeout(() => {
//         dispatch(loginFailureTimeout());
//       }, LOGIN_FAILURE_WAIT_TIME);
//       dispatch(loginFailure(response.response));
//     }
//     );
//   };
// }

function login(hostname, username, password, afterLoginPath) {
  return dispatch => {
    dispatch(loginRequest(hostname, username, password));
    dispatch(setHostAddress(hostname));
    dispatch(loginSuccess(''));
    dispatch(startup(hostname, username, password));
    browserHistory.push(afterLoginPath);
  };
}

export { login };
