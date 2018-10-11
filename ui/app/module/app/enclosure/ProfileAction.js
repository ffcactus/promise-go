import { ActionType } from './ConstValue';

/**
 * This action will be called when user select profile resource.
 */
export function selectResource() {
  return {
    type: ActionType.EP_UI_SELECT_RESOURCE,
  };
}
