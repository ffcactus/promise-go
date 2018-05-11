import {
  ActionType,
  ServerAppState,
} from './ConstValue';
import {
  List,
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
  appState: ServerAppState.LOADING,
  // We need record the default servergroup because of it's special role.
  defaultServerGroupUri: null,
  currentServerGroupUri: null,
  currentServerUri: null,
  currentServer: null,
  serverGroupList: new List(),
  serverList: new List(),
  serverTask: new Map(),
  openCreateServerGroupDialog: false,
  openAddServerDialog: false,
};

export const serverApp = (state = defaultState, action) => {
  let tempDefaultServerGroupUri;
  let tempCurrentServerGroupUri;
  let tempCurrentServerUri;
  let temp;
  switch (action.type) {
    // App
    case ActionType.APP_INIT_START:
      return {
        ...state,
        appState: ServerAppState.LOADING,
        currentServer: null,
        currentServerUri: action.info.currentServerUri,
        currentServerGroupUri: action.info.currentServerGroupUri,
      };
    case ActionType.APP_INIT_SUCCESS:
      // Find the default servergroup.
      for (const each of action.info.serverGroupList.Members) {
        if (each.Name === 'all') {
          tempDefaultServerGroupUri = each.URI;
        }
      }
      // Set the current servergroup.
      if (state.currentServerGroupUri === null) {
        tempCurrentServerGroupUri = tempDefaultServerGroupUri;
      }
      if (state.currentServerUri === null) {
        tempCurrentServerUri = action.info.serverServerGroupList.Members.length === 0 ? null : action.info.serverServerGroupList.Members[0].ServerURI;
      }
      return {
        ...state,
        appState: ServerAppState.NORMAL,
        serverGroupList: List(action.info.serverGroupList.Members),
        serverList: new List(action.info.serverServerGroupList.Members.filter((each) => {
          return each.ServerGroupURI === tempCurrentServerGroupUri;
        }).map((each) => {
          return {
            key: each.ServerURI,
            value: null,
          };
        })),
        defaultServerGroupUri: tempDefaultServerGroupUri,
        currentServerGroupUri: tempCurrentServerGroupUri,
        currentServerUri: tempCurrentServerUri,
      };
    case ActionType.APP_INIT_FAILURE:
      return {
        ...state,
        appState: ServerAppState.FAILURE,
        serverGroupList: new List(),
        serverList: new List(),
      };
    case ActionType.APP_EXIT:
      return defaultState;
    // Server
    // Server.REST.Get
    case ActionType.SERVER_REST_GET_START:
      return state;
    case ActionType.SERVER_REST_GET_SUCCESS:
      // It indicates that the server in the list is loaded or the server selected by user is loaded.
      return {
        ...state,
        serverList: state.serverList.map((each) => {
          if (each.key === action.info.URI) {
            return {key: each.key, value: action.info};
          }
          return each;
        }),
        currentServer: action.info.URI === state.currentServerUri ? action.info : state.currentServer,
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
      // It indicates the a server is created, but we won't do anything here. we don't know if the server belongs to the current list.
      // So we just care about the SSG_WS_CREATE
      return state;
    case ActionType.SERVER_WS_UPDATE:
      // If the server in the list.
      return {
        ...state,
        serverList: state.serverList.map((each) => {
          if (each.key === action.info.URI) {
            return {
              key: each.key,
              value: action.info
            };
          }
          return each;
        }),
        currentServer: action.info.URI === state.currentServerUri ? action.info : state.currentServer,
      };
    case ActionType.SERVER_WS_DELETE:
    case ActionType.SERVER_WS_DELETE_LIST:
      return state;
    // Server.UI
    // Server.UI.List
    case ActionType.SERVER_UI_LIST_SELECT:
      return {
        ...state,
        currentServerUri: action.info,
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
      return {
        ...state,
        serverGroupList: List(action.info.Members),
      };
    case ActionType.SG_REST_GETLIST_MESSAGE:
    case ActionType.SG_REST_GETLIST_EXCEPTION:
      return Object.assign({}, state, {
        serverGroupList: new List(),
      });
    // ServerGroup.WS
    case ActionType.SG_WS_CREATE:
      return {
        ...state,
        serverGroupList: state.serverGroupList.push(action.info)
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
        serverGroupList: [state.defaultServerGroupUri]
      };
    // ServerGroup.UI
    // ServerGroup.UI.List
    case ActionType.SG_UI_LIST_SELECT:
      return {
        ...state,
        currentServerGroupUri: action.info,
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
        serverList: new List(),
        currentServerUri: null,
        currentServer: null,
      };
    case ActionType.SSG_REST_GETLIST_SUCCESS:
      // This indicate that the server list can be updated, we need also select the current server.
      temp = new List(action.info.Members.map((each) => {
        return {
          key: each.ServerURI,
          value: null
        };
      }));
      return {
        ...state,
        serverList: temp,
        currentServerUri: temp.size === 0 ? null : temp.get(0).key,
      };
    case ActionType.SSG_REST_GETLIST_MESSAGE:
    case ActionType.SSG_REST_GETLIST_EXCEPTION:
      // report error.
      // TODO
      return state;
    // Server-ServerGroup.WS
    // We need check if the server belongs to the group selected.
    case ActionType.SSG_WS_CREATE:
      if (action.info.ServerGroupURI === state.currentServerGroupUri) {
        return {
          ...state,
          serverList: state.serverList.push({
            key: action.info.ServerURI,
            value: null,
          }),
        };
      }
      return state;
    case ActionType.SSG_WS_UPDATE:
      return state;
    case ActionType.SSG_WS_DELETE:
      if (action.info.ServerGroupURI === state.currentServerGroupUri) {
        return {
          ...state,
          serverList: state.serverList.filter((each) => each.key === action.info.ServerURI),
        };
      }
      return state;
    // Task.WS
    case ActionType.TASK_WS_CREATE:
    case ActionType.TASK_WS_UPDATE:
      // If the task related to the server in the list.
      for (let i = 0; i < state.serverList.size; i++) {
        if (state.serverList.get(i).key === action.info.TargetURI) {
          return {
            ...state,
            serverTask: state.serverTask.set(action.info.TargetURI, action.info),
          };
        }
      }
      return state;
    case ActionType.TASK_WS_DELETE:
      return state;
    // Others
    default:
      return state;
  }
};
