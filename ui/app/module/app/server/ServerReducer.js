import { ActionType, ServerAppState } from './ConstValue';

const defaultState = {
  state: ServerAppState.UNKNOWN,
  defaultServerGroup: {},
  currentServerGroup: 'all',
  serverGroupList: [],
  serverList: [],
};

const server = (state = defaultState, action) => {
  let tempDefaultServerGroup;
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
      for (const each of action.info.Members) {
        if (each.Name === 'all') {
          tempDefaultServerGroup = each;
        }
      }
      return {
        ...state,
        serverGroupList: action.info.Members.map((each) => {
          return {
            URI: each.URI,
            Name: each.Name,
          };
        }),
        defaultServerGroup: tempDefaultServerGroup
      };
    case ActionType.GET_SERVERGROUP_LIST_FAILURE:
      return Object.assign({}, state, {
        serverGroupList: [],
      });
    case ActionType.ON_SERVERGROUP_CREATE:
      return {
        ...state,
        serverGroupList: state.serverGroupList.concat(action.info)
      };
    case ActionType.ON_SERVER_SERVERGROUP_UPDATE:
      return {
        ...state,
        serverGroupList: state.serverGroupList.map((each) => {
          if (each.ID === action.info.ID) {
            return action.info;
          }
          return each;
        })
      };
    case ActionType.ON_SERVERGROUP_DELETE:
      return {
        ...state,
        serverGroupList: state.serverGroupList.filter(each => each.ID !== action.info)
      };
    case ActionType.ON_SERVERGROUP_DELETE_COLLECTION:
      return {
        ...state,
        serverGroupList: [state.defaultServerGroup]
      };
    default:
      return state;
  }
};

export default server;
