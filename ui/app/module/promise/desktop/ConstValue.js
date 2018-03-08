export const DesktopState = Object.freeze(
  {
    'UNKNOWN': 'UNKNOWN',
    'INIT_START': 'INIT_START',
    'UNSERVICEABLE': 'UNSERVICEABLE',
    'CONNECTION_ERROR': 'CONNECTION_ERROR',
    'APP_COLLECTION': 'APP_COLLECTION',
  }
);

export const DesktopAction = Object.freeze(
  {
    'SET_INIT_START': 'SET_INIT_START',
    'SET_INIT_FAILURE': 'SET_INIT_FAILURE',
    'SET_APP_COLLECTION': 'SET_APP_COLLECTION',
  }
);
