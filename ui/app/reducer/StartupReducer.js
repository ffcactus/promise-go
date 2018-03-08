import * as types from '../actions/types';

const defaultStartupState = {
  ws: {
    state: 'unknown',
  }
};

const startup = (state = defaultStartupState, action) => {
  switch (action.type) {
    case types.WS_NOT_SUPPORT:
    case types.WS_SERVER_CONNECTING:
    case types.WS_SERVER_CONNECTED:
    case types.WS_SERVER_CONNECT_FAILED:
    case types.WS_SERVER_CONNECTION_CLOSED:
      return Object.assign({}, state, {
        ws: {
          state: action.type
        }
      });
    default:
      return state;
  }
};

export default startup;
