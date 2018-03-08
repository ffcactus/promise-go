import { connect } from 'react-redux';
import PrivateRouteContainer from './PrivateRouteContainer';
import { LoginState } from '../login/ConstValue';
const PrivateRoute = connect(state => {
  isAuthenticated: state.session.state === LoginState.LOGGED;
})(PrivateRouteContainer);

export default PrivateRoute;
