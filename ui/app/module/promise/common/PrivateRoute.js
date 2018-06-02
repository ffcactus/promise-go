import React from 'react';
import PropTypes from 'prop-types';
import { Redirect } from 'react-router-dom';
import { Route } from 'react-router';
import { connect } from 'react-redux';
import { LoginState } from '../login/ConstValue';

class PrivateRouteContainer extends React.Component {
  render() {
    const {
      isAuthenticated,
      component: Component,
      ...rest
    } = this.props;

    // render is a function which will return the component.
    return (
      <Route
        {...rest}
        render={p =>
          isAuthenticated
            ? <Component {...this.props} />
            : (
              <Redirect to={{
                pathname: '/login',
                state: { from: p.location }
              }} />
            )
        }
      />
    );
  }
}

PrivateRouteContainer.propTypes = {
  isAuthenticated: PropTypes.bool,
  component: PropTypes.func,
};

const PrivateRoute = connect(state => ({
  isAuthenticated: state.session.state === LoginState.LOGGED
}))(PrivateRouteContainer);

export default PrivateRoute;

