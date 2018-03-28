import * as Client from './Client';
import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';
import * as ServerAction from './ServerAction';
import * as ServerGroupAction from './ServerGroupAction';

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

function appInitFailure() {
  return {
    type: ActionType.APP_INIT_FAILURE,
  };
}

export function appInit(hostname) {
  WsAction.registerMessageAction('Server', ServerAction.onServerMessage);
  WsAction.registerMessageAction('ServerGroup', ServerGroupAction.onServerGroupMessage);
  return (dispatch, state) => {
    dispatch(appInitStart());
    // First, we need get all the servergroup.
    dispatch(ServerGroupAction.getServerGroupListStart());
    Client.getServerGroupList(hostname).then((sgResp) => {
      if (sgResp.status === 200) {
        dispatch(ServerGroupAction.getServerGroupListSuccess(sgResp.response));
        // Then we get all the servers from the current servergroup.
        dispatch(ServerAction.getServerListStart());
        Client.getServerList(hostname, state.currentGroup).then((sResp) => {
          if (sResp.status === 200) {
            dispatch(ServerAction.getServerListSuccess(sResp.response));
            dispatch(appInitSuccess());
            return;
          }
          // If status code error in getting server list, init fails.
          dispatch(ServerAction.getServerListFailure());
          dispatch(appInitFailure());
        }).catch((e) => {
          // if exception raised in getting server list, init fails.
          dispatch(ServerAction.getServerListFailure(e));
          dispatch(appInitFailure(e));
        });
        return;
      }
      // if status code error in getting servergroup list, init fails.
      dispatch(ServerGroupAction.getServerGroupListFailure());
      dispatch(appInitFailure());
    }).catch((e) => {
      // if exception raised in getting servergroup list, init fails.
      dispatch(ServerGroupAction.getServerGroupListFailure(e));
      dispatch(appInitFailure(e));
    });
  };
}