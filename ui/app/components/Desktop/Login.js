import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import { login } from '../../actions/LoginAction';
import Styles from '../../styles/login.css';
import { LoginState } from '../../reducer/SessionReducer';

class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      hostname: window.location.hostname,
      username: '',
      password: ''
    };
    this.handleHostnameChange = this.handleHostnameChange.bind(this);
    this.handleUsernameChange = this.handleUsernameChange.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleHostnameChange(event) {
    this.setState({ hostname: event.target.value });
  }

  handleUsernameChange(event) {
    this.setState({ username: event.target.value });
  }

  handlePasswordChange(event) {
    this.setState({ password: event.target.value });
  }

  handleSubmit(event) {
    event.preventDefault();

    // If we can't find a next path after login, we go to root.
    const nextPathname = this.props.location.state ? this.props.location.state.nextPathname : '/';
    this.props.dispatch(login(this.state.hostname, this.state.username, this.state.password, nextPathname));
  }


  render() {
    const isLoginButtonDisabled = () => {
      switch (this.props.session.state) {
        case LoginState.LOGGING:
        case LoginState.LOGIN_FAILURE_WAIT:
          return true;
        default:
          return false;
      }
    };

    const isErrorMessageShow = () => {
      return (this.props.session.state === LoginState.LOGIN_FAILURE_WAIT);
    };

    const getLoginFailureDescription = () => {
      const info = this.props.session.loginFailureInfo;
      return info ? info.description : null;
    };

    return (
      <div styleName="loginForm">
        <form id="login" onSubmit={this.handleSubmit}>
          <p styleName="loginTitle">Promise</p>
          <section styleName="loginInput">
            <input id="hostname" type="text" placeholder="hostname" onChange={this.handleHostnameChange} />
            <input id="username" type="text" placeholder="username" onChange={this.handleUsernameChange} />
            <input id="password" type="password" placeholder="password" onChange={this.handlePasswordChange} />
          </section>
          <section styleName="loginSubmit">
            <input type="submit" value="login" disabled={isLoginButtonDisabled()} />
          </section>
          <section styleName="loginFailureMessage">
            <p styleName={isErrorMessageShow() ? 'showError' : 'hideError'}>{getLoginFailureDescription()}</p>
          </section>
        </form>
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { session } = state;
  return { session };
}

Login.propTypes = {
  session: PropTypes.object,
  location: PropTypes.object,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(CSSModules(Login, Styles));
