import * as types from '../actions/types';

const defaultSettingState = {
  upgradeBundle: {}
};

const setting = (state = defaultSettingState, action) => {
  switch (action.type) {
    case types.SETTING_SELECT_UPGRADE_BUNDLE:
      return Object.assign({}, state, {
        upgradeBundle: action.info.files[0]
      });
    default:
      return state;
  }
};

export default setting;
