import * as types from './types';

export function settingUploadUpgradeBundleAction(fileName) {
  return {
    type: types.SETTING_SELECT_UPGRADE_BUNDLE,
    info: fileName
  };
}
