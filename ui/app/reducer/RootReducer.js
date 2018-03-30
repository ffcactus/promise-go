import { routerReducer as routing } from 'react-router-redux';
import { combineReducers } from 'redux';
import desktop from '../module/promise/desktop/DesktopReducer';
import session from '../module/promise/login/SessionReducer';
import serverApp from '../module/app/server/ServerReducer';


const rootReducer = combineReducers({
  desktop,
  global,
  session,
  serverApp,
  routing
});

export default rootReducer;
