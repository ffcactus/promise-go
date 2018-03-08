import * as ActionConst from '../constValue/action/Desktop';
import * as StateConst from '../constValue/state/Desktop';

const defaultState = {
  hostname: null,
  state: StateConst.UNKNOWN,
  serviceState: null
};

const desktop = (state = defaultState, action) => {
  switch (action.type) {
    case ActionConst.DESKTOP_INIT_START:
      return Object.assign({}, state, {
        hostname: action.info.hostname,
        state: StateConst.GET_GLOBAL_START
      });
    case ActionConst.DESKTOP_INIT_SUCCESS:
      return Object.assign({}, state, {
        state: StateConst.GET_GLOBAL_SUCCESS
      });
    case ActionConst.DESKTOP_INIT_FAILURE:
      return Object.assign({}, state, {
        state: StateConst.GET_GLOBAL_FAILURE
      });
    default:
      return state;
  }
};

export default desktop;
