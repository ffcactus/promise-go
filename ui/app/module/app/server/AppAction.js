import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';
import * as TaskAction from './TaskAction';
import * as ServerAction from './ServerAction';
import * as ServerGroupAction from './ServerGroupAction';
import * as ServerServerGroupAction from './ServerServerGroupAction';
import { doGet } from '../../promise/common/Client';
const Promise = require('promise');

/**
 * Init the server App.
 * We can pass 2 parameters:
 * ServerGroupID and ServerID, if they are passed, you need select them in the UI.
 * Otherwise, we need get the servergroup list, and the servers belongs to the default servergroup.
 * @param {string} servergroup The ID of the servergroup to be select.
 * @param {string} server The ID of the server to be select.
 */
export function appInit(servergroup, server) {
  WsAction.registerMessageAction('Server', ServerAction.onServerMessage);
  WsAction.registerMessageAction('ServerGroup', ServerGroupAction.onServerGroupMessage);
  WsAction.registerMessageAction('ServerServerGroup', ServerServerGroupAction.onServerServerGroupMessage);
  WsAction.registerMessageAction('Task', TaskAction.onTaskMessage);
  return (dispatch, getState) => {
    const prefix = 'http://' + getState().session.hostname;
    dispatch({
      type: ActionType.APP_SERVER_INIT_START,
      info: {
        currentServerUri: server ? '/promise/v1/server/' + server : null,
        currentServerGroupUri: servergroup ? '/promise/v1/servergroup/' + servergroup : null,
      }
    });

    const sURL = prefix + '/promise/v1/server';
    const sgURL = prefix + '/promise/v1/servergroup';
    const ssgURL = prefix + '/promise/v1/server-servergroup';
    Promise.all([
      doGet(sURL),
      doGet(sgURL),
      doGet(ssgURL),
    ]).then((responses) => {
      if (responses[0].status === 200 && responses[1].status === 200 && responses[1].status === 200) {
        dispatch({
          type: ActionType.APP_SERVER_INIT_SUCCESS,
          info: {
            serverList: responses[0].response,
            serverGroupList: responses[1].response,
            serverServerGroupList: responses[2].response,
          }
        });
        return;
      }
      // if not HTTP status 200.
      dispatch({
        type: ActionType.APP_SERVER_INIT_FAILURE,
      });
    }).catch((e) => {
      dispatch({
        type: ActionType.APP_SERVER_INIT_FAILURE,
        info: e,
      });
    });
  };
}
