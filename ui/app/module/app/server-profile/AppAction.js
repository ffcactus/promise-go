import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';
import * as ServerProfileAction from './ServerProfileAction';
import { doGet } from '../../promise/common/Client';
const Promise = require('promise');

export function appInit(model, config) {
  WsAction.registerMessageAction('ServerProfile', ServerProfileAction.onMessage);
  return (dispatch, getState) => {
    const prefix = 'http://' + getState().session.hostname;
    const modelURL = prefix + '/promise/v1/adaptermodel';
    dispatch({
      type: ActionType.APP_INIT_START,
      info: {
        model,
        config
      }
    });
    Promise.all([
      doGet(modelURL)
    ]).then((responses) => {
      if(responses[0].status === 200) {
        dispatch({
          type: ActionType.APP_INIT_SUCCESS
        });
        return;
      }
      dispatch({
        type: ActionType.APP_INIT_FAILURE
      });
    }).catch(e => {
      dispatch({
        type: ActionType.APP_INIT_FAILURE,
        info: e,
      });
    });
  };
}
