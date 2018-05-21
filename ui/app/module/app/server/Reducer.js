import { ActionType, ServerAppState, ServerDetailState } from './ConstValue';
import { List, Map, Set } from 'immutable';

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
  serverDetailState: ServerDetailState.EMPTY,
  // We need record the default servergroup because of it's special role.
  defaultServerGroupUri: null,  // The default servergroup, the one named 'all'.
  currentServerGroupUri: null,  // The servergroup selected.
  currentServerUri: null,       // The server selected.
  currentServer: null,          // The server shows in detail.
  currentServerSet: new Set(),    // To indicate the server that should be put in the list.
  serverGroupList: new List(),  // To record all the servergroups.
  ssgSet: null,            // To record all the SSGs.
  serverList: new List(),       // To record all the servers
  serverTask: new Map(),        // To record all the tasks whos target is the servers in the list.
  openCreateServerGroupDialog: false,
  openAddServerDialog: false,
  serverExist: false,
  serverGroupExist: false,
  serverIndex: undefined,
  serverGroupIndex: undefined,
};

export const serverApp = (state = defaultState, action) => {
  let tempDefaultServerGroupUri;
  let tempCurrentServerGroupUri;
  let tempCurrentServerUri;
  let tempCurrentServerSet;
  let tempServerGroupExist;
  let tempServerExist;
  let tempServerIndex;
  let arraylength;
  switch (action.type) {
    // App
    case ActionType.APP_INIT_START:
      return {
        ...state,
        appState: ServerAppState.LOADING,
        defaultServerGroupUri: null,
        currentServer: null,
        currentServerSet: new Set(),
        serverGroupList: new List(),
        ssgSet: null,
        serverList: new List(),
        serverTask: new Map(),
        openCreateServerGroupDialog: false,
        openAddServerDialog: false,
        currentServerUri: action.info.currentServerUri,
        currentServerGroupUri: action.info.currentServerGroupUri,
      };
    case ActionType.APP_INIT_SUCCESS:
      tempCurrentServerUri = state.currentServerUri;
      tempServerExist = false;
      tempServerGroupExist = false;
      // Find the default servergroup.
      for (const each of action.info.serverGroupList.Members) {
        if (each.Name === 'all') {
          tempDefaultServerGroupUri = each.URI;
        }
        if (each.URI === state.currentServerGroupUri) {
          tempServerGroupExist = true;
        }
      }
      arraylength = action.info.serverList.Members.length;
      for (let i = 0; i < arraylength; i++) {
        if (action.info.serverList.Members[i].URI === state.currentServerUri) {
          tempServerIndex = i;
          tempServerExist = true;
        }
      }
      // Set the current servergroup.
      if (state.currentServerGroupUri === null) {
        tempServerGroupExist = true;
        tempCurrentServerGroupUri = tempDefaultServerGroupUri;
      }
      if (state.currentServerUri === null) {
        tempServerExist = true;
        tempCurrentServerUri = action.info.serverServerGroupList.Members.length === 0 ? null : action.info.serverServerGroupList.Members[0].ServerURI;
      }
      return {
        ...state,
        appState: ServerAppState.NORMAL,
        defaultServerGroupUri: tempDefaultServerGroupUri,
        currentServerGroupUri: tempCurrentServerGroupUri,
        currentServerUri: tempCurrentServerUri,
        currentServerSet: Set(action.info.serverServerGroupList.Members.filter((each) => {
          return each.ServerGroupURI === tempCurrentServerGroupUri;
        }).map(each => each.ServerURI)),
        serverGroupList: List(action.info.serverGroupList.Members),
        ssgSet: action.info.serverServerGroupList.Members,
        serverList: List(action.info.serverList.Members),
        serverExist: tempServerExist,
        serverGroupExist: tempServerGroupExist,
        serverIndex: tempServerIndex,
      };
    case ActionType.APP_INIT_FAILURE:
      return {
        ...state,
        appState: ServerAppState.FAILURE,
        serverDetailState: ServerDetailState.FAILURE,
      };
    case ActionType.APP_EXIT:
      return defaultState;
    // Server
    // Server.REST.Get
    case ActionType.SERVER_REST_GET_START:
      return {
        ...state,
        currentServer: null,
        serverDetailState: ServerDetailState.LOADING,
      };
    case ActionType.SERVER_REST_GET_SUCCESS:
      return {
        ...state,
        currentServer: action.info,
        serverDetailState: ServerDetailState.NORMAL,
      };
    case ActionType.SERVER_REST_GET_MESSAGE:
    case ActionType.SERVER_REST_GET_EXCEPTION:
      return {
        ...state,
        serverDetailState: ServerDetailState.FAILURE
      };
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
      return {
        ...state,
        serverList: state.serverList.push(action.info),
      };
    case ActionType.SERVER_WS_UPDATE:
      // If the server in the list.
      return {
        ...state,
        serverList: state.serverList.map((each) => {
          if (each.ID === action.info.ID) {
            return action.info;
          }
          return each;
        }),
        currentServer: action.info.URI === state.currentServerUri ? action.info : state.currentServer,
      };
    case ActionType.SERVER_WS_DELETE:
      return {
        ...state,
        serverList: state.serverList.filter((each) => each.ID !== action.info.ID),
        currentServer: action.info.URI === state.currentServerUri ? null : state.currentServer,
      };
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
    case ActionType.SG_REST_CREATE_SUCCESS:
    case ActionType.SG_REST_CREATE_MESSAGE:
    case ActionType.SG_REST_CREATE_EXCEPTION:
      return state;
    // ServerGroup.REST.GetList
    case ActionType.SG_REST_GETLIST_START:
    case ActionType.SG_REST_GETLIST_SUCCESS:
    case ActionType.SG_REST_GETLIST_MESSAGE:
    case ActionType.SG_REST_GETLIST_EXCEPTION:
      return state;
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
      tempCurrentServerSet = [];
      arraylength = state.ssgSet.length;
      for (let i = 0; i < arraylength; i++) {
        if (state.ssgSet[i].ServerGroupURI === action.info) {
          tempCurrentServerSet.push(state.ssgSet[i].ServerURI);
        }
      }
      return {
        ...state,
        currentServer: null,
        currentServerUri: null,
        currentServerGroupUri: action.info,
        currentServerSet: new Set(tempCurrentServerSet),
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
    case ActionType.SSG_REST_GETLIST_SUCCESS:
    case ActionType.SSG_REST_GETLIST_MESSAGE:
    case ActionType.SSG_REST_GETLIST_EXCEPTION:
      return state;
    // Server-ServerGroup.WS
    // We need check if the server belongs to the group selected.
    case ActionType.SSG_WS_CREATE:
      return {
        ...state,
        ssgSet: state.ssgSet.concat(action.info),
        currentServerSet: action.info.ServerGroupURI === state.currentServerGroupUri ? state.currentServerSet.add(action.info.ServerURI) : state.currentServerSet,
      };
    case ActionType.SSG_WS_UPDATE:
      return state;
    case ActionType.SSG_WS_DELETE:
      return {
        ...state,
        ssgSet: state.ssgSet.filter((each) => each.ID !== action.info.ID),
        currentServerSet: action.info.ServerGroupURI === state.currentServerGroupUri ? state.currentServerSet.delete(action.info.ServerURI) : state.currentServerSet,
      };
    // Task.WS
    case ActionType.TASK_WS_CREATE:
    case ActionType.TASK_WS_UPDATE:
      // If the task related to the server in the list.
      for (let i = 0; i < state.serverList.size; i++) {
        if (state.serverList.get(i).URI === action.info.TargetURI) {
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
