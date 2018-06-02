import { ActionType } from './ConstValue';
import { createPostAction } from '../../promise/common/Client';

export function addServer(name) {
  return (dispatch, getState) => {
    const request = {
      Name: name,
      Hostname: 'Mock',
      Username: 'Username',
      Password: 'Password'
    };

    createPostAction(
      '/promise/v1/server/action/discover',
      request,
      ActionType.SERVER_REST_ADD_START,
      ActionType.SERVER_REST_ADD_SUCCESS,
      ActionType.SERVER_REST_ADD_MESSAGE,
      ActionType.SERVER_REST_ADD_EXCEPTION,
    )(dispatch, getState);
  };
}
