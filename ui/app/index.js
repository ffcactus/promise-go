import 'babel-polyfill';
import React from 'react';
import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';
import { Route, Switch } from 'react-router';
import { Provider } from 'react-redux';
import { ConnectedRouter } from 'react-router-redux';
import { render } from 'react-dom';
import createHistory from 'history/createBrowserHistory';
import { routerMiddleware } from 'react-router-redux';
import rootReducer from './RootReducer';
import Login from './module/promise/login/Login';
import PrivateRoute from './module/promise/common/PrivateRoute';
import DesktopContainer from './module/promise/desktop/DesktopContainer';
import Settings from './module/app/settings/Settings';
import Animation from './module/app/animation/Animation';
import ServerContainer from './module/app/server/ServerContainer';

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const history = createHistory();
const router = routerMiddleware(history);
const store = createStore(rootReducer, composeEnhancers(applyMiddleware(router, thunk)));

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <Switch>
        <Route path="/login" component={Login} />
        <PrivateRoute exact path="/" component={DesktopContainer} />
        <PrivateRoute path="/app/settings" component={Settings} />
        <PrivateRoute path="/app/animation" component={Animation} />
        <PrivateRoute appName="Server" path="/app/server" hostname={window.location.hostname} component={ServerContainer} />
      </Switch>
    </ConnectedRouter>
  </Provider>,
  document.getElementById('root')
);
