import { ActionType, LoginState } from './ConstValue';

const defaultSessionState = {
  hostname: null,
  state: LoginState.LOGOUT,
  username: null,
  token: null,
  loginFailureInfo: null
};


const session = (state = defaultSessionState, action) => {
  switch (action.type) {
    case ActionType.LOGIN_REQUEST:
      return {
        ...state,
        state: LoginState.LOGGING,
        token: null
      };
    case ActionType.LOGIN_SUCCESS:
      return {
        ...state,
        state: LoginState.LOGGED,
        username: action.info.username,
        hostname: action.info.hostname,
        token: action.token
      };
    case ActionType.LOGIN_FAILURE:
      return {
        ...state,
        state: LoginState.LOGIN_FAILURE_WAIT,
        username: null,
        token: null,
        loginFailureInfo: action.info
      };
    case ActionType.LOGIN_FAILURE_TIMEOUT:
      return {
        ...state,
        state: LoginState.LOGOUT,
        username: null,
        token: null
      };
    case ActionType.LOGOUT_REQUEST:
      return state;
    case ActionType.LOGOUT_SUCCESS:
      return defaultSessionState;
    case ActionType.LOGOUT_FAILURE:
      return state;
    default:
      return state;
  }
};

export default session;

