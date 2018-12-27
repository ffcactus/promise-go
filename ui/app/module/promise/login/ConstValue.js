export const LoginState = Object.freeze(
  {
    'LOGGING': 'LOGGING',
    'LOGGED': 'LOGGED',
    'LOGOUT': 'LOGOUT',
    'LOGIN_FAILURE_WAIT': 'LOGIN_FAILURE_WAIT',
  }
);

export const ActionType = Object.freeze(
  {
    'LOGIN_START': 'LOGIN_START',
    'LOGIN_FAILURE': 'LOGIN_FAILURE',
    'LOGIN_FAILURE_TIMEOUT': 'LOGIN_FAILURE_TIMEOUT',
    'LOGIN_SUCCESS': 'LOGIN_SUCCESS',
    'GOOGLE_LOGIN_SUCCESS': 'GOOGLE_LOGIN_SUCCESS',
    'LOGOUT_REQUEST': 'LOGOUT_REQUEST',
    'LOGOUT_FAILURE': 'LOGOUT_FAILURE',
    'LOGOUT_SUCCESS': 'LOGOUT_SUCCESS',
  }
);
