import 'babel-polyfill';
import React from 'react';
import { createStore, applyMiddleware, compose } from 'redux';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import { ConnectedRouter } from 'react-router-redux';
import { render } from 'react-dom';
import createHistory from 'history/createBrowserHistory';
import { routerMiddleware } from 'react-router-redux';
import rootReducer from './RootReducer';
import Desktop from './module/promise/desktop/Desktop';

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const history = createHistory();
const router = routerMiddleware(history);
const store = createStore(rootReducer, composeEnhancers(applyMiddleware(router, thunk)));

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <Desktop />
    </ConnectedRouter>
  </Provider>,
  document.body,
  // document.getElementById('root')
);
