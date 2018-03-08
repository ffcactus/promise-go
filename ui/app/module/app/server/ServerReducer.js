import { ActionType, ServerAppState } from './ConstValue';

const defaultState = {
  state: ServerAppState.UNKNOWN,
  groupList: [{
    Name: 'All'
  }, {
    Name: 'Rack'
  }, {
    Name: 'Blade'
  }],
  serverList: [],
};

const server = (state = defaultState, action) => {
  switch(action.type) {
    case ActionType.APP_INIT_START:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_START,
        serverList: [],
      });
    case ActionType.APP_INIT_SUCCESS:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_SUCCESS,
        serverList: [],
      });
    case ActionType.APP_INIT_FAILURE:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_FAILURE,
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
        state: ServerAppState.APP_INIT_FAILURE,
        serverList: [],
      });
    default:
      return state;
  }
};

export default server;
