import * as types from './types';

export function setHostAddress(hostAddress) {
  return {
    type: types.OPEN,
    info: { hostAddress }
  };
}
