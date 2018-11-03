import {
  ActionType,
  AppState,
  EnclosureResource,
  ResourceDetailState,
} from './ConstValue';
import { List, Map, OrderedMap } from 'immutable';

const defaultState = {
  appState: AppState.LOADING,
  currentResource: EnclosureResource.Enclosure,
  enclosureOrderedMap: new OrderedMap(),
  // records the enclosures latest task, key is enclosure URI.
  enclosureTask: new Map(),
  poolList: new List(),
  profileList: new List(),
  taskMap: new Map(),
  enclosureUri: null,
  enclosure: null,
  resourceDetailState: ResourceDetailState.EMPTY,
  openDiscoverEnclosureDialog: false,
};

export const enclosureApp = (state = defaultState, action) => {
  let tempEnclosure = null;

  switch (action.type) {
    // App
    case ActionType.APP_ENCLOSURE_INIT_START:
      return {
        ...state,
        appState: AppState.LOADING,
        enclosureUri: action.info.enclosureUri ? action.info.enclosureUri : null,
      };
    case ActionType.APP_ENCLOSURE_INIT_SUCCESS:
      return {
        ...state,
        appState: AppState.NORMAL,
        enclosureOrderedMap: OrderedMap(action.info.enclosureList.Members.map(each => {
          return [each.URI, {...each, UI: null}];
        })),
      };
    case ActionType.APP_ENCLOSURE_INIT_FAILURE:
      return {
        ...state,
        appState: AppState.FAILURE,
      };
    case ActionType.APP_ENCLOSURE_EXIT:
      return defaultState;
    // Enclosure
    // Enclosure.REST
    // Enclosure.REST.Get
    case ActionType.ENCLOSURE_REST_GET_START:
      return {
        ...state,
        enclosure: null,
        resourceDetailState: ResourceDetailState.LOADING
      };
    case ActionType.ENCLSOURE_REST_GET_SUCCESS:
      return {
        ...state,
        enclosure: action.info,
        resourceDetailState: ResourceDetailState.READY
      };
    case ActionType.ENCLSOURE_REST_GET_EXCEPTION:
    case ActionType.ENCLSOURE_REST_GET_MESSAGE:
      return {
        ...state,
        resourceDetailState: ResourceDetailState.FAILURE
      };
    // Enclosure.REST.Discover
    case ActionType.ENCLOSURE_REST_DISCOVER_START:
      return state;
    case ActionType.ENCLOSURE_REST_DISCOVER_SUCCESS:
      return {
        ...state,
        openDiscoverEnclosureDialog: false,
      };
    case ActionType.ENCLOSURE_REST_DISCOVER_MESSAGE:
      return state;
    case ActionType.ENCLOSURE_REST_DISCOVER_EXCEPTION:
      return state;
    // Enclosure.WS
    case ActionType.ENCLOSURE_WS_CREATE:
    case ActionType.ENCLOSURE_WS_UPDATE:
      tempEnclosure = state.enclosureOrderedMap.get(action.info.URI);
      return {
        ...state,
        enclosureOrderedMap: state.enclosureOrderedMap.set(action.info.URI, {
          ID: action.info.ID,
          URI: action.info.URI,
          Category: action.info.Category,
          Name: action.info.Name,
          State: action.info.State,
          Health: action.info.Health,
          UI: tempEnclosure === undefined ? null : tempEnclosure.UI,
        }),
        enclosure: action.info.URI === state.enclosureUri ? action.info : state.enclosure,
      };
    case ActionType.ENCLOSURE_WS_DELETE:
      return {
        ...state,
        // TODO
        enclosureOrderedMap: state.enclosureOrderedMap.delete(action.info.ID),
        enclosure: action.info.URI === state.enclosureUri ? null : state.enclosure,
      };
    case ActionType.ENCLOSURE_WS_DELETE_LIST:
      return state;
    // Enclosure.UI
    // Enclosure.UI.Resource
    case ActionType.ENCLOSURE_UI_SELECT_RESOURCE:
      return {
        ...state,
        currentResource: EnclosureResource.Enclosure
      };
    // Enclosure.UI.List
    case ActionType.ENCLOSURE_UI_SELECT:
      return {
        ...state,
        enclosureUri: action.info,
      };
    // Enclosure.UI.Dialog
    // Enclosure.UI.Dialog.Discover
    case ActionType.ENCLOSURE_UI_DIALOG_DISCOVER_OPEN:
      return {
        ...state,
        openDiscoverEnclosureDialog: true
      };
    case ActionType.ENCLOSURE_UI_DIALOG_DISCOVER_CLOSE:
      return {
        ...state,
        openDiscoverEnclosureDialog: false,
      };
    // Profile
    // Profile.UI
    // Profile.UI.Select
    case ActionType.EP_UI_SELECT_RESOURCE:
      return {
        ...state,
        currentResource: EnclosureResource.Profile
      };
    // IDPool
    // IDPool.UI
    // IDPool.UI.Select
    case ActionType.IDPOOL_UI_SELECT_RESOURCE:
      return {
        ...state,
        currentResource: EnclosureResource.IDPool
      };
    // Task
    // Task.WS
    case ActionType.TASK_WS_CREATE:
    case ActionType.TASK_WS_UPDATE:
      tempEnclosure = state.enclosureOrderedMap.get(action.info.TargetURI);
      if (tempEnclosure !== undefined) {
        return {
          ...state,
          enclosureOrderedMap: state.enclosureOrderedMap.set(action.info.TargetURI, {
            ...tempEnclosure,
            UI: {
              ...tempEnclosure.UI,
              task: action.info
            }
          }),
        };
      }
      return state;
    default:
      return state;
  }
};
