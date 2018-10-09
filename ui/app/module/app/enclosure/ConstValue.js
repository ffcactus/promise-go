export const AppState = Object.freeze(
  {
    'LOADING': 'LOADING',
    'NORMAL': 'NORMAL',
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
    // Enclosure.UI
    'ENCLOSURE_UI_SELECT': 'ENCLOSURE_UI_SELECT',

    // Profile
    // Profile.UI
    'EP_UI_SELECT': 'EP_UI_SELECT',

    // IDPool
    // IDPool.UI
    'IDPOOL_UI_SELECT': 'IDPOOL_UI_SELECT',
  }
);
