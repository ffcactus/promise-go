import * as types from '../actions/types';

const defaultServerState = {
  current: null,
  serverList: [],
  serverLoading: false,
  serverListLoading: false,
};

const server = (state = defaultServerState, action) => {
  let newState;
  switch (action.type) {
    case types.SERVER_GET_BEGIN:
      return Object.assign({}, state, {
        current: null,
        serverLoading: true,
      });
    case types.SERVER_GET_SUCCESS:
      return Object.assign({}, state, {
        current: action.info,
        serverLoading: false
      });
    case types.SERVER_GET_FAILURE:
      return Object.assign({}, state, {
        current: null,
        serverLoading: false,
      });
    case types.SERVER_LIST_LOAD_BEGIN:
      return Object.assign({}, state, {
        serverList: [],
        current: null,
        serverListLoading: true
      });
    case types.SERVER_LIST_LOAD_SUCCESS:
      return Object.assign({}, state, {
        serverList: action.info.Members.map((each) => {
          return {
            URI: each.URI,
            Name: each.Name,
            State: each.State,
            Health: each.Health
          };
        }),
        serverListLoading: false,
      });
    case types.SERVER_LIST_LOAD_FAILURE:
      return Object.assign({}, state, {
        serverList: [],
        current: null,
        serverListLoading: false,
      });
    case types.SERVER_CREATE:
      newState = Object.assign({}, state, {
        current: state.current,
        serverListLoading: state.serverListLoading
      });
      newState.serverList.push({
        Name: action.info.Name,
        Uri: action.info.Uri,
        State: action.info.State,
        Health: action.info.Health
      });
      return newState;
    case types.SERVER_UPDATE:
      newState = Object.assign({}, state, {});
      newState.serverList.map((each) => {
        if (action.info.Uri === each.Uri) {
          each.Name = action.info.Name;
          each.Uri = action.info.Uri;
          each.State = action.info.State;
          each.Health = action.info.Health;
        }
      });
      if (newState.current !== null && newState.current.Id === action.info.Id) {
        newState.current = action.info;
      }
      return newState;
    default:
      return state;
  }
};

export default server;
