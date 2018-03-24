import { ActionType, ServerAppState } from './ConstValue';

const defaultState = {
  state: ServerAppState.UNKNOWN,
  serverGroupList: [],
  serverList: [],
};

const server = (state = defaultState, action) => {
  switch (action.type) {
    case ActionType.APP_INIT_START:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_START,
        serverGroupList: [],
        serverList: [],
      });
    case ActionType.APP_INIT_SUCCESS:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_SUCCESS,
      });
    case ActionType.APP_INIT_FAILURE:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_FAILURE,
        serverGroupList: [],
        serverList: [],
      });
    case ActionType.GET_SERVER_LIST_START:
      return state;
    case ActionType.GET_SERVER_LIST_SUCCESS:
      return Object.assign({}, state, {
        serverList: action.info.Members.map((each) => {
          return {
            URI: each.URI,
            Name: each.Name,
            State: each.State,
            Health: each.Health
          };
        })
      });
    case ActionType.GET_SERVER_LIST_FAILURE:
      return Object.assign({}, state, {
        serverList: [],
      });
    case ActionType.GET_SERVERGROUP_LIST_START:
      return state;
    case ActionType.GET_SERVERGROUP_LIST_SUCCESS:
      return Object.assign({}, state, {
        serverGroupList: action.info.Members.map((each) => {
          return {
            URI: each.URI,
            Name: each.Name,
          };
        })
      });
    case ActionType.GET_SERVERGROUP_LIST_FAILURE:
      return Object.assign({}, state, {
        serverGroupList: [],
      });
    default:
      return state;
  }
};

export default server;
