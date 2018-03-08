import * as types from '../actions/types';

const LoginState = {
  LOGGING: 'logging',
  LOGGED: 'logged',
  LOGOUT: 'logout',
  LOGIN_FAILURE_WAIT: 'loginFailureWait'
};

const defaultSessionState = {
  state: LoginState.LOGOUT,
  username: null,
  token: null,
  loginFailureInfo: null
};


const session = (state = defaultSessionState, action) => {
  switch (action.type) {
    case types.LOGIN_REQUEST:
      return {
        state: LoginState.LOGGING,
        username: action.username,
        token: null
      };
    case types.LOGIN_SUCCESS:
      return {
        state: LoginState.LOGGED,
        username: null,
        token: action.token
      };
    case types.LOGIN_FAILURE:
      return {
        state: LoginState.LOGIN_FAILURE_WAIT,
        username: null,
        token: null,
        loginFailureInfo: action.info
      };
    case types.LOGIN_FAILURE_TIMEOUT:
      return {
        state: LoginState.LOGOUT,
        username: null,
        token: null
      };
    case types.LOGOUT_REQUEST:
      return state;
    case types.LOGOUT_SUCCESS:
      return defaultSessionState;
    case types.LOGOUT_FAILURE:
      return state;
    default:
      return state;
  }
};

export { session, LoginState };

