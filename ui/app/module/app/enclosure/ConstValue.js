export const AppState = Object.freeze(
  {
    'LOADING': 'LOADING',
    'NORMAL': 'NORMAL',
    'FAILURE': 'FAILURE',
  }
);

export const ResourceDetailState = Object.freeze(
  {
    'EMPTY': 'EMPTY',
    'LOADING': 'LOADING',
    'READY': 'READY',
    'FAILURE': 'FAILURE',
  }
);

export const EnclosureResource = Object.freeze(
  {
    'Enclosure': 'Enclosure',
    'Group': 'Group',
    'Profile': 'Profile',
    'IDPool': 'IDPool',
  }
);

export const ActionType = Object.freeze(
  {
    // App
    'APP_ENCLOSURE_INIT_START': 'APP_ENCLOSURE_INIT_START',
    'APP_ENCLOSURE_INIT_SUCCESS': 'APP_ENCLOSURE_INIT_SUCCESS',
    'APP_ENCLOSURE_INIT_FAILURE': 'APP_ENCLOSURE_INIT_FAILURE',
    'APP_ENCLOSURE_EXIT': 'APP_ENCLOSURE_EXIT',

    // Enclosure
    // Enclosure.REST
    // Enclosure.REST.Get
    'ENCLOSURE_REST_GET_START': 'ENCLOSURE_REST_GET_START',
    'ENCLSOURE_REST_GET_SUCCESS': 'ENCLSOURE_REST_GET_SUCCESS',
    'ENCLSOURE_REST_GET_MESSAGE': 'ENCLSOURE_REST_GET_MESSAGE',
    'ENCLSOURE_REST_GET_EXCEPTION': 'ENCLSOURE_REST_GET_EXCEPTION',
    // Enclosure.REST.Discover
    'ENCLOSURE_REST_DISCOVER_START': 'ENCLOSURE_REST_DISCOVER_START',
    'ENCLOSURE_REST_DISCOVER_SUCCESS': 'ENCLOSURE_REST_DISCOVER_SUCCESS',
    'ENCLOSURE_REST_DISCOVER_MESSAGE': 'ENCLOSURE_REST_DISCOVER_MESSAGE',
    'ENCLOSURE_REST_DISCOVER_EXCEPTION': 'ENCLOSURE_REST_DISCOVER_EXCEPTION',

    // Enclosure.UI
    // Enclosure.UI.Resource
    'ENCLOSURE_UI_SELECT_RESOURCE': 'ENCLOSURE_UI_SELECT_RESOURCE',
    // Enclosure.UI.List
    'ENCLOSURE_UI_SELECT': 'ENCLOSURE_UI_SELECT',
    // Enclosure.UI.Discover
    'ENCLOSURE_UI_DIALOG_DISCOVER_OPEN': 'ENCLOSURE_UI_DIALOG_DISCOVER_OPEN',
    'ENCLOSURE_UI_DIALOG_DISCOVER_CLOSE': 'ENCLOSURE_UI_DIALOG_DISCOVER_CLOSE',

    // Profile
    // Profile.UI
    'EP_UI_SELECT_RESOURCE': 'EP_UI_SELECT_RESOURCE',

    // IDPool
    // IDPool.UI
    'IDPOOL_UI_SELECT_RESOURCE': 'IDPOOL_UI_SELECT_RESOURCE',

    // Task.WS
    'TASK_WS_CREATE': 'TASK_WS_CREATE',
    'TASK_WS_UPDATE': 'TASK_WS_UPDATE',
    'TASK_WS_DELETE': 'TASK_WS_DELETE',
  }
);
