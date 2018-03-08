import { DesktopAction, DesktopState } from './ConstValue';

const defaultState = {
  hostname: null,
  state: DesktopState.UNKNOWN,
  serviceState: null
};

const desktop = (state = defaultState, action) => {
  switch (action.type) {
    case DesktopAction.SET_INIT_START:
      return Object.assign({}, state, {
        hostname: action.info.hostname,
        state: DesktopState.INIT_START
      });
    case DesktopAction.SET_INIT_FAILURE:
      return Object.assign({}, state, {
        state: DesktopState.UNSERVICEABLE
      });
    case DesktopAction.SET_APP_COLLECTION:
      return Object.assign({}, state, {
        state: DesktopState.APP_COLLECTION
      });
    default:
      return state;
  }
};

export default desktop;
