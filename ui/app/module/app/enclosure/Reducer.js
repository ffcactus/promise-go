import {
  ActionType,
  AppState,
  EnclosureResource,
} from './ConstValue';
import { List, Map } from 'immutable';

const defaultState = {
  appState: AppState.LOADING,
  currentResource: EnclosureResource.Enclosure,
  enclosureList: new List(),
  poolList: new List(),
  profileList: new List(),
  taskMap: new Map(),
  enclosureUri: null,
  enclosure: null,
};

export const enclosureApp = (state = defaultState, action) => {
  switch (action.type) {
    // App
    case ActionType.APP_ENCLOSURE_INIT_START:
      return {
        ...state,
        appState: AppState.LOADING,
        enclosureList: new List(),
        poolList: new List(),
        profileList: new List(),
        taskMap: new Map(),
        enclosureUri: action.info.enclosureUri,
      };
    case ActionType.APP_ENCLOSURE_INIT_SUCCESS:
      return {
        ...state,
        appState: AppState.NORMAL,
        enclosureList: List(action.info.enclosureList.Members),
      };
    case ActionType.APP_ENCLOSURE_INIT_FAILURE:
      return {
        ...state,
        appState: AppState.FAILURE,
      };
    case ActionType.APP_ENCLOSURE_EXIT:
      return defaultState;
    // Enclosure
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
    default:
      return state;
  }
};
