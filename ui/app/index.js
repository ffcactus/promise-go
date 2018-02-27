import 'babel-polyfill';
import React from 'react';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import { Route, Switch } from 'react-router';
import { Provider } from 'react-redux';
import { ConnectedRouter } from 'react-router-redux';
import { render } from 'react-dom';
import createHistory from 'history/createBrowserHistory';
import { routerMiddleware } from 'react-router-redux';
import rootReducer from './reducer/RootReducer';
import Login from './module/promise/login/Login';
import PrivateRoute from './module/promise/common/PrivateRoute';
import DesktopContainer from './module/promise/desktop/DesktopContainer';
import Settings from './module/app/settings/Settings';
import Phone from './module/app/phone/Phone';
import Server from './module/app/server/Server';
const history = createHistory();
const router = routerMiddleware(history);
const store = createStore(rootReducer, applyMiddleware(router, thunk));

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <Switch>
        <Route path="/login" component={Login} />
        <PrivateRoute exact path="/" component={DesktopContainer} />
        <PrivateRoute path="/app/settings" component={Settings} />
        <PrivateRoute path="/app/phone" component={Phone} />
        <PrivateRoute path="/app/server" component={Server} />
      </Switch>
    </ConnectedRouter>
  </Provider>,
  document.getElementById('root')
);
