import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';
import * as ServerAction from './ServerAction';
import * as ServerGroupAction from './ServerGroupAction';
import * as ServerServerGroupAction from './ServerServerGroupAction';

function appInitStart() {
  return {
    type: ActionType.APP_INIT_START,
  };
}

function appInitSuccess() {
  return {
    type: ActionType.APP_INIT_SUCCESS,
  };
}

// function appInitFailure() {
//   return {
//     type: ActionType.APP_INIT_FAILURE,
//   };
// }

// export function appInit() {
//   WsAction.registerMessageAction('Server', ServerAction.onServerMessage);
//   WsAction.registerMessageAction('ServerGroup', ServerGroupAction.onServerGroupMessage);
//   WsAction.registerMessageAction('ServerServerGroup', ServerServerGroupAction.onServerServerGroupMessage);
//   return (dispatch, getState) => {
//     const hostname = getState().session.hostname;
//     dispatch(appInitStart());
//     // First, we need get all the servergroup.
//     dispatch(ServerGroupAction.getServerGroupListStart());
//     Client.getServerGroupList(hostname).then((sgResp) => {
//       if (sgResp.status === 200) {
//         dispatch(ServerGroupAction.getServerGroupListSuccess(sgResp.response));
//         dispatch(appInitSuccess());
//         return;
//       }
//       // if status code error in getting servergroup list, init fails.
//       dispatch(ServerGroupAction.getServerGroupListFailure());
//       dispatch(appInitFailure());
//     }).catch((e) => {
//       // if exception raised in getting servergroup list, init fails.
//       dispatch(ServerGroupAction.getServerGroupListFailure(e));
//       dispatch(appInitFailure(e));
//     });
//   };
// }

/**
 * Init the server App.
 *
 */
export function appInit() {
  WsAction.registerMessageAction('Server', ServerAction.onServerMessage);
  WsAction.registerMessageAction('ServerGroup', ServerGroupAction.onServerGroupMessage);
  WsAction.registerMessageAction('ServerServerGroup', ServerServerGroupAction.onServerServerGroupMessage);
  return (dispatch) => {
    dispatch(appInitStart());
    dispatch(appInitSuccess());
  };
}
