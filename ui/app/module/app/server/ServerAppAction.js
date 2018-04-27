import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';
import * as ServerAction from './ServerAction';
import * as ServerGroupAction from './ServerGroupAction';
import * as ServerServerGroupAction from './ServerServerGroupAction';
import { doGet } from '../../../client/common';

/**
 * Init the server App.
 * We need get the default servergroup URI in initialization.
 */
export function appInit() {
  WsAction.registerMessageAction('Server', ServerAction.onServerMessage);
  WsAction.registerMessageAction('ServerGroup', ServerGroupAction.onServerGroupMessage);
  WsAction.registerMessageAction('ServerServerGroup', ServerServerGroupAction.onServerServerGroupMessage);
  return (dispatch, getState) => {
    dispatch({
      type: ActionType.APP_INIT_START,
    });
    const uri = getState().session.hostname + '/promise/v1/servergroup?/$filter=Name eq \'all\'';
    doGet(uri).then((resp)=> {
      if (resp.status === 200) {
        dispatch({
          type: ActionType.APP_INIT_SUCCESS,
          info: resp.Members[0].URI,
        });
        return;
      }
      dispatch({
        type: ActionType.APP_INIT_FAILURE,
      });
    }).catch(() => {
      dispatch({
        type: ActionType.APP_INIT_FAILURE,
      });
    });
  };
}
