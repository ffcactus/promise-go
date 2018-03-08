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
      ...props
    } = this.props;

    // render is a function which will return the component.
    return (
      <Route
        {...props}
        render={p =>
          isAuthenticated
            ? <Component {...p} />
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

