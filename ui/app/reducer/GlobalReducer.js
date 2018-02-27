import * as types from '../actions/types';

const defaultSettingState = {
  hostAddress: null
};

const global = (state = defaultSettingState, action) => {
  switch (action.type) {
    case types.OPEN:
      return Object.assign({}, state, {
        hostAddress: action.info.hostAddress
      });
    default:
      return state;
  }
};

export default global;
