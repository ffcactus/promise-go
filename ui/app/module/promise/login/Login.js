import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import { login, googleLogin } from './LoginAction';
import { LoginState } from './ConstValue';
import FullSizeDiv from '../common/Widget/FullSizeDiv';
import GoogleLogin from 'react-google-login';
import Styles from './login.css';

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
    this.onGoogleLoginFailure = this.onGoogleLoginFailure.bind(this);
    this.onGoogleLoginSuccess = this.onGoogleLoginSuccess.bind(this);
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

  onGoogleLoginSuccess(userObject) {
    const from = this.props.location.state ? this.props.location.state.from : '/';
    this.props.dispatch(googleLogin(this.state.hostname, userObject, from));
  }

  onGoogleLoginFailure(reason) {
    console.info(reason);
  }

  handleSubmit(event) {
    event.preventDefault();

    // If we can't find a next path after login, we go to root.
    const from = this.props.location.state ? this.props.location.state.from : '/';
    this.props.dispatch(login(this.state.hostname, this.state.username, this.state.password, from));
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
      return info && info.Description ? info.Description : 'Unknown Error';
    };

    return (
      <div>
        <FullSizeDiv>
          <div styleName="loginForm">
            <form onSubmit={this.handleSubmit}>
              <p styleName="loginTitle">Promise</p>
              <section styleName="loginInput">
                <input type="text" value={this.state.hostname} placeholder="hostname" onChange={this.handleHostnameChange} />
                <input type="text" placeholder="username" onChange={this.handleUsernameChange} />
                <input type="password" placeholder="password" onChange={this.handlePasswordChange} />
              </section>
              <section styleName="loginSubmit">
                <input type="submit" value="login" disabled={isLoginButtonDisabled()} />
              </section>
              <section styleName="loginSubmit">
                <GoogleLogin
                  clientId="988618725491-umsr9vm0m439hijt990q500nci03bjv5.apps.googleusercontent.com"
                  buttonText="Login"
                  onSuccess={this.onGoogleLoginSuccess}
                  onFailure={this.onGoogleLoginFailure}
                />
              </section>
              <section styleName="loginFailureMessage">
                <p styleName={isErrorMessageShow() ? 'showError' : 'hideError'}>{getLoginFailureDescription()}</p>
              </section>
            </form>
          </div>
        </FullSizeDiv>
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
