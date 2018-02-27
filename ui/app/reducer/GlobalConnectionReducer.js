import * as types from '../actions/types';

const SVC_STATE_UNKNOWN = 'SVC_STATE_UNKNOWN';
const SVC_STATE_OK = 'SVC_STATE_OK';
const SVC_STATE_WARNING = 'SVC_STATE_WARNING';
const SVC_STATE_CRITICAL = 'SVC_STATE_CRITICAL';

const defaultSettingState = {
  serverServiceState: SVC_STATE_UNKNOWN,
  taskServiceState: SVC_STATE_UNKNOWN,
};

const globalConnection = (state = defaultSettingState, action) => {
  switch (action.type) {
    // Server.
    case types.INIT_SERVER_SVC_START:
      return Object.assign({}, state, {
        serverServiceState: SVC_STATE_UNKNOWN
      });
    case types.INIT_SERVER_SVC_SUCCESS:
      return Object.assign({}, state, {
        serverServiceState: SVC_STATE_OK
      });
    case types.INIT_SERVER_SVC_FAILURE:
      return Object.assign({}, state, {
        serverServiceState: SVC_STATE_CRITICAL
      });
    case types.SERVER_SVC_CLOSE:
      return Object.assign({}, state, {
        serverServiceState: SVC_STATE_WARNING
      });
    // Task.
    case types.INIT_TASK_SVC_START:
      return Object.assign({}, state, {
        taskServiceState: SVC_STATE_UNKNOWN
      });
    case types.INIT_TASK_SVC_SUCCESS:
      return Object.assign({}, state, {
        taskServiceState: SVC_STATE_OK
      });
    case types.INIT_TASK_SVC_FAILURE:
      return Object.assign({}, state, {
        taskServiceState: SVC_STATE_CRITICAL
      });
    case types.TASK_SVC_CLOSE:
      return Object.assign({}, state, {
        taskServiceState: SVC_STATE_WARNING
      });
    default:
      return state;
  }
};

export default globalConnection;
