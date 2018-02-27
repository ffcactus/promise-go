import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { ConnectedRouter } from 'react-router-redux';
import { connect, Provider } from 'react-redux';
import CSSModules from 'react-css-modules';
import RTRouter from './RTRouter';
import styles from '../styles/main.css';
// styleName="rootBackground"
class Root extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    const { store, history } = this.props;
    const { getState } = store;
    return (
      <Provider store={store}>
        <ConnectedRouter history={history}>
          <div id="fullscreen" >
            <RTRouter history={history} getState={getState} />
          </div>
        </ConnectedRouter>
      </Provider>
    );
  }
}

function mapStateToProps(state) {
  const { startup } = state;
  return { startup };
}

Root.propTypes = {
  store: PropTypes.object.isRequired,
  dispatch: PropTypes.func,
  history: PropTypes.object.isRequired,
  onStartup: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(Root, styles));
