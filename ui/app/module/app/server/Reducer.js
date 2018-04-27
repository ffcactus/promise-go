import {
  ActionType,
  ServerAppState
} from './ConstValue';
import {
  Map
} from 'immutable';

/**
 * The store updated in the this way:
 * When the App will mount, get all the servergroups.
 * When the App is mount, set the selected servergroup to be the default one.
 * When the a servergroup is selected, get all the servers belong to this servergroup.
 * The servers will be put into a Map, the URI will be the key and the value will be the GetServerResponse.
 * After the server list is loaded, set the selected server to be the one on the top.
 * When an server element in the list is mounted, get the server and put the response as the value to the server map.
 * When a servergroup is added, we add it to servergroup list.
 * When a servergroup is removed, we remove it from servergroup list, and if it's the one been selected,
 * we selected the default one.
 * When a server is added to a servergroup and the servergroup is the one been selected, we add it to server map.
 * When a server is removed from a servergroup and the servergroup is the one been selected, we remove it from server map,
 * and if it's the server been selected, we choose another one unless it's empty.
 *
 */

const defaultState = {
  appState: ServerAppState.APP_INIT_START,
  // We need record the default servergroup because of it's special role.
  defaultServerGroup: null,
  currentServerGroup: null,
  currentServer: null,
  serverDetail: {},
  serverGroupList: [],
  serverList: new Map(),
  openCreateServerGroupDialog: false,
  openAddServerDialog: false,
};

const serverApp = (state = defaultState, action) => {
  let tempDefaultServerGroup;
  let tempSelectedServer;
  switch (action.type) {
    // App
    case ActionType.APP_INIT_START:
      return defaultState;
    case ActionType.APP_INIT_SUCCESS:
      return {
        ...state,
        appState: ServerAppState.APP_INIT_SUCCESS,
        defaultServerGroup: action.info,
      };
    case ActionType.APP_INIT_FAILURE:
      return{
        ...state,
        appState: ServerAppState.APP_INIT_FAILURE,
        serverGroupList: [],
        serverList: new Map(),
      };
    case ActionType.APP_EXIT:
      return defaultState;
    // Server
    // Server.REST.Get
    case ActionType.SERVER_REST_GET_START:
      return state;
    case ActionType.SERVER_REST_GET_SUCCESS:
      return {
        ...state,
        serverList: state.serverList.set(action.info.URI, action.info),
      };
    case ActionType.SERVER_REST_GET_MESSAGE:
      return state;
    case ActionType.SERVER_REST_GET_EXCEPTION:
      return state;
    // Server.REST.GetList
    case ActionType.SERVER_REST_GETLIST_START:
    case ActionType.SERVER_REST_GETLIST_SUCCESS:
    case ActionType.SERVER_REST_GETLIST_MESSAGE:
    case ActionType.SERVER_REST_GETLIST_EXCEPTION:
      return state;
    // Server.WS
    case ActionType.SERVER_WS_CREATE:
    case ActionType.SERVER_WS_UPDATE:
    case ActionType.SERVER_WS_DELETE:
    case ActionType.SERVER_WS_DELETE_LIST:
      return state;
    // Server.UI
    // Server.UI.List
    case ActionType.SERVER_UI_LIST_SELECT:
      return {
        ...state,
        currentServer: action.info,
      };
    // Server.UI.Dialog
    case ActionType.SERVER_UI_DIALOG_ADD_OPEN:
    case ActionType.SERVER_UI_DIALOG_ADD_CLOSE:
      return state;

    // ServerGroup
    // ServerGroup.REST.Create
    case ActionType.SG_REST_CREATE_START:
      return state;
    case ActionType.SG_REST_CREATE_SUCCESS:
      return state;
    case ActionType.SG_REST_CREATE_MESSAGE:
    case ActionType.SG_REST_CREATE_EXCEPTION:
      return state;
    // ServerGroup.REST.GetList
    case ActionType.SG_REST_GETLIST_START:
      return state;
    case ActionType.SG_REST_GETLIST_SUCCESS:
      for (const each of action.info.Members) {
        if (each.Name === 'all') {
          tempDefaultServerGroup = each.URI;
        }
      }
      return {
        ...state,
        serverGroupList: action.info.Members,
        defaultServerGroup: tempDefaultServerGroup,
        // If no servergroup is selected, select the default one.
        currentServerGroup: (state.currentServerGroup === null) ? tempDefaultServerGroup : state.currentServerGroup,
      };
    case ActionType.SG_REST_GETLIST_MESSAGE:
    case ActionType.SG_REST_GETLIST_EXCEPTION:
      return Object.assign({}, state, {
        serverGroupList: [],
      });
    // ServerGroup.WS
    case ActionType.SG_WS_CREATE:
      return {
        ...state,
        serverGroupList: state.serverGroupList.concat(action.info)
      };
    case ActionType.SG_WS_UPDATE:
      return {
        ...state,
        serverGroupList: state.serverGroupList.map((each) => {
          if (each.ID === action.info.ID) {
            return action.info;
          }
          return each;
        })
      };
    case ActionType.SG_WS_DELETE:
      return {
        ...state,
        serverGroupList: state.serverGroupList.filter(each => each.ID !== action.info.ID)
      };
    case ActionType.SG_WS_DELETE_LIST:
      return {
        ...state,
        serverGroupList: [state.defaultServerGroup]
      };
    // ServerGroup.UI
    // ServerGroup.UI.List
    case ActionType.SG_UI_LIST_SELECT:
      return {
        ...state,
        currentServerGroup: action.info,
      };
    // ServerGroup.UI.Dialog
    case ActionType.SG_UI_DIALOG_ADD_OPEN:
      return {
        ...state,
        openCreateServerGroupDialog: true,
        openAddServerDialog: false
      };
    case ActionType.SG_UI_DIALOG_ADD_CLOSE:
      return {
        ...state,
        openCreateServerGroupDialog: false,
      };

    // Server-ServerGroup.REST
    // Server-ServerGroup.REST.GetList
    case ActionType.SSG_REST_GETLIST_START:
      // clean the server list.
      return {
        ...state,
        serverList: new Map(),
      };
    case ActionType.SSG_REST_GETLIST_SUCCESS:
    // create the server list.
      tempSelectedServer = action.info.Members.length === 0 ? null : action.info.Members[0].ServerURI;
      return {
        ...state,
        // create a map for server list, the key is URI, the value is empty.
        serverList: new Map(action.info.Members.map((each) => {
          return [each.ServerURI, {}];
        })),
        currentServer: state.currentServer === null ? tempSelectedServer : state.currentServer,
      };
    case ActionType.SSG_REST_GETLIST_MESSAGE:
    case ActionType.SSG_REST_GETLIST_EXCEPTION:
      // report error.
      // TODO
      return state;
    // Server-ServerGroup.WS
      // We need check if the server belongs to the group selected.
    case ActionType.SSG_WS_CREATE:
      if (action.info.ServerGroupID === state.currentServerGroup.ID) {
        return {
          ...state,
          serverList: state.serverList.set(action.info.ServerURI, {})
        };
      }
      return state;
    case ActionType.SSG_WS_UPDATE:
      return state;
    case ActionType.SSG_WS_DELETE:
      if (action.info.ServerGroupID === state.currentServerGroup.ID) {
        return {
          ...state,
          serverList: state.serverList.delete(action.info.ServerURI)
        };
      }
      return state;
    // Others
    default:
      return state;
  }
};

export default serverApp;
