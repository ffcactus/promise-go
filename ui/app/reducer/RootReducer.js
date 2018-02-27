import { routerReducer as routing } from 'react-router-redux';
import { combineReducers } from 'redux';
import desktop from '../module/promise/desktop/DesktopReducer';
import startup from './StartupReducer';
import global from './GlobalReducer';
import session from '../module/promise/login/SessionReducer';
import task from './TaskReducer';
import server from './ServerReducer';
import setting from './SettingReducer';
import platform from './PlatformReducer';

const rootReducer = combineReducers({
  desktop,
  platform,
  startup,
  global,
  session,
  task,
  server,
  setting,
  routing
});

export default rootReducer;
