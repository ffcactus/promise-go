import {
  ActionType,
  AppState,
  EnclosureResource,
} from './ConstValue';

const defaultState = {
  appState: AppState.LOADING,
  currentResource: EnclosureResource.Enclosure,
};

export const enclosureApp = (state = defaultState, action) => {
  switch (action.type) {
    // App
    case ActionType.APP_ENCLOSURE_INIT_START:
      return {
        ...state,
        appState: AppState.LOADING,
      };
    case ActionType.APP_ENCLOSURE_INIT_SUCCESS:
      return {
        ...state,
        appState: AppState.NORMAL,
      };
    case ActionType.APP_ENCLOSURE_INIT_FAILURE:
      return {
        ...state,
        appState: AppState.FAILURE,
      };
    case ActionType.APP_ENCLOSURE_EXIT:
      return defaultState;
    // Enclosure.UI.Select
    case ActionType.ENCLOSURE_UI_SELECT:
      return {
        ...state,
        currentResource: EnclosureResource.Enclosure
      };
    // Profile.UI.Select
    case ActionType.EP_UI_SELECT:
      return {
        ...state,
        currentResource: EnclosureResource.Profile
      };
    // Profile.UI.Select
    case ActionType.IDPOOL_UI_SELECT:
      return {
        ...state,
        currentResource: EnclosureResource.IDPool
      };
    default:
      return state;
  }
};
