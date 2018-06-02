import { ActionType } from './ConstValue';

const defaultState = {
  response: null,
  message: null,
  exception: null,
};

export const serverControllerApp = (state = defaultState, action) => {
  switch(action.type) {
    case ActionType.SERVER_REST_ADD_START:
      return defaultState;
    case ActionType.SERVER_REST_ADD_SUCCESS:
      return {
        ...state,
        response: action.info
      };
    case ActionType.SERVER_REST_ADD_MESSAGE:
      return {
        ...state,
        message: action.info
      };
    case ActionType.SERVER_REST_ADD_EXCEPTION:
      return {
        ...state,
        exception: action.info
      };
    default:
      return state;
  }
};
