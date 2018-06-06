import { AppState, ActionType } from './ConstValue';
import { List } from 'immutable';

const defaultState = {
  appState: AppState.LOADING,
};

export const serverProfileApp = (state = defaultState, action) => {
  switch (action.type) {
    case ActionType.APP_INIT_START:
      return {
        ...state,
        appState: AppState.LOADING,
        modelList: new List(),
        ConfigList: new List()
      };
    default:
      return state;
  }
};
