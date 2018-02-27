import { ActionType, LoginState } from './ConstValue';

const defaultSessionState = {
  state: LoginState.LOGOUT,
  username: null,
  token: null,
  loginFailureInfo: null
};


const session = (state = defaultSessionState, action) => {
  switch (action.type) {
    case ActionType.LOGIN_REQUEST:
      return {
        state: LoginState.LOGGING,
        username: action.username,
        token: null
      };
    case ActionType.LOGIN_SUCCESS:
      return {
        state: LoginState.LOGGED,
        username: null,
        token: action.token
      };
    case ActionType.LOGIN_FAILURE:
      return {
        state: LoginState.LOGIN_FAILURE_WAIT,
        username: null,
        token: null,
        loginFailureInfo: action.info
      };
    case ActionType.LOGIN_FAILURE_TIMEOUT:
      return {
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

