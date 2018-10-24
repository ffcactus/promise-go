import { ActionType } from './ConstValue';
import * as WsAction from '../../promise/ws/WsAction';
import * as TaskAction from './TaskAction';
import * as EnclosureAction from './EnclosureAction';
import { doGet } from '../../promise/common/Client';
const Promise = require('promise');

/**
 * Init the enclosure App.
 * We can let user pre-select resources before open this App by specifiy their ID
 * in the parameters.
 * @param {string} enclosure The pre-selected enclosure ID.
 * @param {string} profile The pre-selected profile ID.
 * @param {string} pool The pre-selected pool ID.
 */
export function appInit(enclosure, profile, pool) {
  WsAction.registerMessageAction('Enclosure', EnclosureAction.onMessage);
  WsAction.registerMessageAction('Task', TaskAction.onTaskMessage);
  return (dispatch, getState) => {
    const prefix = 'http://' + getState().session.hostname;
    dispatch({
      type: ActionType.APP_ENCLOSURE_INIT_START,
      info: {
        selectedEnclosure: enclosure ? enclosure : null,
        selectedProfile: profile ? profile : null,
        selectedPool: pool ? pool : null,
      }
    });

    const enclosureURL = prefix + '/promise/v1/enclosure';
    Promise.all([doGet(enclosureURL)]).then((response) => {
      if (response[0].status === 200) {
        dispatch({
          type: ActionType.APP_ENCLOSURE_INIT_SUCCESS,
          info: {
            enclosureList: response[0].response,
          }
        });
        return;
      }
      // if not HTTP status 200.
      dispatch({
        type: ActionType.APP_ENCLOSURE_INIT_FAILURE,
      });
    }).catch((e) => {
      dispatch({
        type: ActionType.APP_ENCLOSURE_INIT_FAILURE,
        info: e,
      });
    });
  };
}
