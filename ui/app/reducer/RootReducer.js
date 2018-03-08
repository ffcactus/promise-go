import { routerReducer as routing } from 'react-router-redux';
import { combineReducers } from 'redux';
import desktop from '../module/promise/desktop/DesktopReducer';
import session from '../module/promise/login/SessionReducer';
import server from '../module/app/server/ServerReducer';


const rootReducer = combineReducers({
  desktop,
  global,
  session,
  server,
  routing
});

export default rootReducer;
