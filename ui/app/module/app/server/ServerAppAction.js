import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';
import * as TaskAction from './TaskAction';
import * as ServerAction from './ServerAction';
import * as ServerGroupAction from './ServerGroupAction';
import * as ServerServerGroupAction from './ServerServerGroupAction';
import { doGet } from '../../../client/common';


/**
 * Init the server App.
 * We can pass 2 parameters:
 * ServerGroupID and ServerID, if they are passed, you need select them in the UI.
 * Otherwise, we need get the servergroup list, and the servers belongs to the default servergroup.
 * @param {string} currentServerGroup The ID of the servergroup to be select.
 * @param {string} currentServer The ID of the server to be select.
 */
export function appInit(currentServerGroup, currentServer) {
  WsAction.registerMessageAction('Server', ServerAction.onServerMessage);
  WsAction.registerMessageAction('ServerGroup', ServerGroupAction.onServerGroupMessage);
  WsAction.registerMessageAction('ServerServerGroup', ServerServerGroupAction.onServerServerGroupMessage);
  WsAction.registerMessageAction('Task', TaskAction.onTaskMessage);
  return (dispatch, getState) => {
    dispatch({
      type: ActionType.APP_INIT_START,
      info: {
        currentServer: null,
        currentServerGroup: null,
      }
    });
    const prefix = 'http://' + getState().session.hostname;
    const sgURI = prefix + '/promise/v1/servergroup';
    const ssgURI = prefix + '/promise/v1/server-servergroup';
    doGet(sgURI).then((sgResp) => {
      if (sgResp.status === 200) {
        return sgResp.response;
      }
      dispatch({
        type: ActionType.APP_INIT_FAILURE,
      });
      return null;
    }).then((servergroup) => {
      doGet(ssgURI).then((ssgResp)=> {
        if (ssgResp.status === 200) {
          dispatch({
            type: ActionType.APP_INIT_SUCCESS,
            info: {
              serverGroupList: servergroup,
              serverServerGroupList: ssgResp.response
            }
          });
          return;
        }
        dispatch({
          type: ActionType.APP_INIT_FAILURE,
        });
      });
    }).catch((e) => {
      dispatch({
        type: ActionType.APP_INIT_FAILURE,
        info: e,
      });
    });
  };
}
