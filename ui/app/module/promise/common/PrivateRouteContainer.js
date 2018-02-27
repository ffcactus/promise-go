import React from 'react';
import PropTypes from 'prop-types';
import { Redirect } from 'react-router-dom';
import { Route } from 'react-router';

class PrivateRouteContainer extends React.Component {
  render() {
    const {
      isAuthenticated,
      component: Component,
      ...props
    } = this.props;

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

export default PrivateRouteContainer;
