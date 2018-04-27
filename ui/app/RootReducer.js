import { routerReducer as routing } from 'react-router-redux';
import { combineReducers } from 'redux';
import desktop from './module/promise/desktop/DesktopReducer';
import session from './module/promise/login/Reducer';
import serverApp from './module/app/server/Reducer';


const rootReducer = combineReducers({
  desktop,
  session,
  serverApp,
  routing
});

export default rootReducer;
