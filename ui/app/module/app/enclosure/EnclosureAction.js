import { ActionType } from './ConstValue';

/**
 * This action will be called when user select enclosure resource..
 */
export function selectResource() {
  return {
    type: ActionType.ENCLOSURE_UI_SELECT_RESOURCE,
  };
}

/**
 * This action means user select a single enclosure.
 * @param {string} uri The URI of the enclosure.
 */
export function selectElement(uri) {
  return {
    type: ActionType.ENCLOSURE_UI_SELECT,
    info: uri
  };
}
