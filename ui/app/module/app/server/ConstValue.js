export const ServerAppState = Object.freeze(
  {
    'APP_INIT_START': 'APP_INIT_START',
    'APP_INIT_SUCCESS': 'APP_INIT_SUCCESS',
    'APP_INIT_FAILURE': 'APP_INIT_FAILURE',
  }
);

export const ActionType = Object.freeze(
  {
    // App
    'APP_INIT_START': 'APP_INIT_START',
    'APP_INIT_SUCCESS': 'APP_INIT_SUCCESS',
    'APP_INIT_FAILURE': 'APP_INIT_FAILURE',
    'APP_EXIT': 'APP_EXIT',
    // Server
    // Server.REST
    // Server.REST.Get
    'SERVER_REST_GET_START': 'SERVER_REST_GET_START',
    'SERVER_REST_GET_SUCCESS': 'SERVER_REST_GET_SUCCESS',
    'SERVER_REST_GET_MESSAGE': 'SERVER_REST_GET_MESSAGE',
    'SERVER_REST_GET_EXCEPTION': 'SERVER_REST_GET_EXCEPTION',
    // Server.REST.GetList
    'SERVER_REST_GETLIST_START': 'SERVER_REST_GETLIST_START',
    'SERVER_REST_GETLIST_SUCCESS': 'SERVER_REST_GETLIST_SUCCESS',
    'SERVER_REST_GETLIST_MESSAGE': 'SERVER_REST_GETLIST_MESSAGE',
    'SERVER_REST_GETLIST_EXCEPTION': 'SERVER_REST_GETLIST_EXCEPTION',
    // Server.WS
    'SERVER_WS_CREATE': 'SERVER_WS_CREATE',
    'SERVER_WS_UPDATE': 'SERVER_WS_UPDATE',
    'SERVER_WS_DELETE': 'SERVER_WS_DELETE',
    'SERVER_WS_DELETE_LIST': 'SERVER_WS_DELETE_LIST',
    // Server.UI
    // Server.UI.List
    'SERVER_UI_LIST_SELECT': 'SERVER_UI_LIST_SELECT',
    // Server.UI.Dialog
    // Server.UI.Dialog.Add
    'SERVER_UI_DIALOG_ADD_OPEN': 'SERVER_UI_DIALOG_ADD_OPEN',
    'SERVER_UI_DIALOG_ADD_CLOSE': 'SERVER_UI_DIALOG_ADD_CLOSE',

    // ServerGroup
    // ServerGroup.REST
    // ServerGroup.REST.Create
    'SG_REST_CREATE_START': 'SG_REST_CREATE_START',
    'SG_REST_CREATE_SUCCESS': 'SG_REST_CREATE_SUCCESS',
    'SG_REST_CREATE_MESSAGE': 'SG_REST_CREATE_MESSAGE',
    'SG_REST_CREATE_EXCEPTION': 'SG_REST_CREATE_EXCEPTION',
    // ServerGroup.REST.GetList
    'SG_REST_GETLIST_START': 'SG_REST_GETLIST_START',
    'SG_REST_GETLIST_SUCCESS': 'SG_REST_GETLIST_SUCCESS',
    'SG_REST_GETLIST_MESSAGE': 'SG_REST_GETLIST_MESSAGE',
    'SG_REST_GETLIST_EXCEPTION': 'SG_REST_GETLIST_EXCEPTION',
    // ServerGroup.WS
    'SG_WS_CREATE': 'SG_WS_CREATE',
    'SG_WS_UPDATE': 'SG_WS_UPDATE',
    'SG_WS_DELETE': 'SG_WS_DELETE',
    'SG_WS_DELETE_LIST': 'SG_WS_DELETE_LIST',
    // ServerGroup.UI
    // ServerGroup.UI.List
    'SG_UI_LIST_SELECT': 'SG_UI_LIST_SELECT',
    // ServerGroup.UI.Dialog
    // ServerGroup.UI.Dialog.Add
    'SG_UI_DIALOG_ADD_OPEN': 'SG_UI_DIALOG_ADD_OPEN',
    'SG_UI_DIALOG_ADD_CLOSE': 'SG_UI_DIALOG_ADD_CLOSE',

    // Server-ServerGroup
    // Server-ServerGroup.REST
    // Server-ServerGroup.REST.GetList
    'SSG_REST_GETLIST_START': 'SSG_REST_GETLIST_START',
    'SSG_REST_GETLIST_SUCCESS': 'SSG_REST_GETLIST_SUCCESS',
    'SSG_REST_GETLIST_MESSAGE': 'SSG_REST_GETLIST_MESSAGE',
    'SSG_REST_GETLIST_EXCEPTION': 'SSG_REST_GETLIST_EXCEPTION',
    // Server-ServerGroup.WS
    'SSG_WS_CREATE': 'SSG_WS_CREATE',
    'SSG_WS_UPDATE': 'SSG_WS_UPDATE',
    'SSG_WS_DELETE': 'SSG_WS_DELETE',
    'SSG_WS_DELETE_LIST': 'SSG_WS_DELETE_LIST',
  }
);
