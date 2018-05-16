import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';
import * as TaskAction from './TaskAction';
import * as ServerAction from './ServerAction';
import * as ServerGroupAction from './ServerGroupAction';
import * as ServerServerGroupAction from './ServerServerGroupAction';
import { doGet } from '../../../client/common';
const Promise = require('promise');

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
        currentServerUri: null,
        currentServerGroupUri: null,
      }
    });
    const prefix = 'http://' + getState().session.hostname;
    const sURI = prefix + '/promise/v1/server';
    const sgURI = prefix + '/promise/v1/servergroup';
    const ssgURI = prefix + '/promise/v1/server-servergroup';
    Promise.all([doGet(sURI), doGet(sgURI), doGet(ssgURI)]).then((responses) => {
      if (responses[0].status === 200 && responses[1].status === 200 && responses[1].status === 200) {
        dispatch({
          type: ActionType.APP_INIT_SUCCESS,
          info: {
            serverList: responses[0].response,
            serverGroupList: responses[1].response,
            serverServerGroupList: responses[2].response,
          }
        });
        return;
      }
      dispatch({
        type: ActionType.APP_INIT_FAILURE,
      });
    }).catch((e) => {
      dispatch({
        type: ActionType.APP_INIT_FAILURE,
        info: e,
      });
    });
  };
}
