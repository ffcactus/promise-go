import { ActionType, ServerAppState } from './ConstValue';
import { Map } from 'immutable';

const defaultState = {
  state: ServerAppState.UNKNOWN,
  // We need record the default servergroup because of it's special role.
  defaultServerGroup: {},
  currentServerGroup: {},
  serverGroupList: [],
  serverList: new Map(),
  openCreateServerGroupDialog: false,
  openAddServerDialog: false,
};

const serverApp = (state = defaultState, action) => {
  let tempDefaultServerGroup;
  switch (action.type) {
    case ActionType.APP_INIT_START:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_START,
      });
    case ActionType.APP_INIT_SUCCESS:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_SUCCESS,
      });
    case ActionType.APP_INIT_FAILURE:
      return Object.assign({}, state, {
        state: ServerAppState.APP_INIT_FAILURE,
        serverGroupList: [],
        serverList: new Map(),
      });
    case ActionType.APP_EXIT:
      return defaultState;
    // Get server list.
    case ActionType.GET_SERVER_LIST_START:
      return state;
    case ActionType.GET_SERVER_LIST_SUCCESS:
      return {
        ...state,
        serverList: new Map(action.info.Members.map((each) => {
          return [
            each.ServerURI,
            {}
          ];
        }))
      };
    case ActionType.GET_SERVER_LIST_FAILURE:
      return {
        ...state,
        serverList: new Map(),
      };
    // Get server.
    case ActionType.GET_SERVER_START:
      return state;
    case ActionType.GET_SERVER_SUCCESS:
      return {
        ...state,
        serverList: state.serverList.set(action.info.URI, action.info),
      };
    case ActionType.GET_SERVER_FAILURE:
      return state;
    // servergroup REST.
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
        serverGroupList: action.info.Members,
        defaultServerGroup: tempDefaultServerGroup,
        currentServerGroup: tempDefaultServerGroup
      };
    case ActionType.GET_SERVERGROUP_LIST_FAILURE:
      return Object.assign({}, state, {
        serverGroupList: [],
      });
    case ActionType.CREATE_SERVERGROUP_START:
      return state;
    case ActionType.CREATE_SERVERGROUP_SUCCESS:
      return state;
    case ActionType.CREATE_SERVERGROUP_FAILURE:
      return state;
    // server event.
    case ActionType.ON_SERVER_CREATE:
      return state;
    case ActionType.ON_SERVER_UPDATE:
      return state;
    case ActionType.ON_SERVER_DELETE:
      return state;
    // servergroup event.
    case ActionType.ON_SERVERGROUP_CREATE:
      return {
        ...state,
        serverGroupList: state.serverGroupList.concat(action.info)
      };
    case ActionType.ON_SERVERGROUP_UPDATE:
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
        serverGroupList: state.serverGroupList.filter(each => each.ID !== action.info.ID)
      };
    case ActionType.ON_SERVERGROUP_DELETE_COLLECTION:
      return {
        ...state,
        serverGroupList: [state.defaultServerGroup]
      };
    case ActionType.ON_SERVERGROUP_SELECTED:
      return {
        ...state,
        currentServerGroup: action.info,
      };
    // Server-servergroup event.
    // We need check if the server belongs to the group selected.
    case ActionType.ON_SERVER_SERVERGROUP_CREATE:
      if (action.info.ServerGroupID === state.currentServerGroup.ID) {
        return {
          ...state,
          serverList: state.serverList.set(action.info.ServerURI, {})
        };
      }
      return state;
    case ActionType.ON_SERVER_SERVERGROUP_UPDATE:
      return state;
    case ActionType.ON_SERVER_SERVERGROUP_DELETE:
      if (action.info.ServerGroupID === state.currentServerGroup.ID) {
        return {
          ...state,
          serverList: state.serverList.delete(action.info.ServerURI)
        };
      }
      return state;
    // create server group dialog.
    case ActionType.OPEN_CREATE_SERVERGROUP_DIALOG:
      return {
        ...state,
        openCreateServerGroupDialog: true,
        openAddServerDialog: false
      };
    case ActionType.CLOSE_CREATE_SERVERGROUP_DIALOG:
      return {
        ...state,
        openCreateServerGroupDialog: false,
      };
    // add server dialog.
    case ActionType.OPEN_ADD_SERVER_DIALOG:
      return {
        ...state,
        openAddServerDialog: true,
        openCreateServerGroupDialog: false,
      };
    case ActionType.CLOSE_ADD_SERVER_DIALOG:
      return {
        ...state,
        openAddServerDialog: false,
      };
    default:
      return state;
  }
};

export default serverApp;
