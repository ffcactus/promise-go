import * as types from '../actions/types';

export const SVC_STATE_UNKNOWN = 'SVC_STATE_UNKNOWN';
export const SVC_STATE_STARTING = 'SVC_STATE_STARTING';
export const SVC_STATE_READY = 'SVC_STATE_READY';
export const SVC_STATE_WARNING = 'SVC_STATE_WARNING';
export const SVC_STATE_CRITICAL = 'SVC_STATE_CRITICAL';

const defaultState = {
  serverServiceState: SVC_STATE_UNKNOWN,
  taskServiceState: SVC_STATE_UNKNOWN,
};

const platform = (state = defaultState, action) => {
  switch (action.type) {
    // Server.
    case types.INIT_SERVER_SVC_START:
      return Object.assign({}, state, {
        serverServiceState: SVC_STATE_STARTING,
      });
    case types.INIT_SERVER_SVC_SUCCESS:
      return Object.assign({}, state, {
        serverServiceState: SVC_STATE_READY
      });
    case types.INIT_SERVER_SVC_FAILURE:
      return Object.assign({}, state, {
        serverServiceState: SVC_STATE_CRITICAL
      });
    case types.SERVER_SVC_CLOSE:
      return Object.assign({}, state, {
        serverServiceState: SVC_STATE_CRITICAL
      });
    // Task.
    case types.INIT_TASK_SVC_START:
      return Object.assign({}, state, {
        taskServiceState: SVC_STATE_STARTING
      });
    case types.INIT_TASK_SVC_SUCCESS:
      return Object.assign({}, state, {
        taskServiceState: SVC_STATE_READY
      });
    case types.INIT_TASK_SVC_FAILURE:
      return Object.assign({}, state, {
        taskServiceState: SVC_STATE_CRITICAL
      });
    case types.TASK_SVC_CLOSE:
      return Object.assign({}, state, {
        taskServiceState: SVC_STATE_CRITICAL
      });
    default:
      return state;
  }
};

export default platform;
