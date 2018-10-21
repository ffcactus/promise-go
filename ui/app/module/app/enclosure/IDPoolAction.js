import { ActionType } from './ConstValue';

/**
 * This action will be called when user select ID Pool.
 */
export function selectResource() {
  return {
    type: ActionType.IDPOOL_UI_SELECT_RESOURCE,
  };
}
