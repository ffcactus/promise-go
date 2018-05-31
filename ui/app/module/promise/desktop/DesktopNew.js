import React from 'react';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import { Route, Switch } from 'react-router';
import { TransitionGroup, CSSTransition } from 'react-transition-group';
import PrivateRoute from '../common/PrivateRoute';
import Login from '../login/Login';
import AppList from './AppList';
import Settings from '../../app/settings/Settings';
import Animation from '../../app/animation/Animation';
import ServerContainer from '../../app/server/ServerContainer';
import styles from './Desktop.css';

class Desktop extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const mainContentFade = {
      appear: this.props.styles.MainContentFadeAppear,
      appearActive: this.props.styles.MainContentFadeAppearActive,
      enter: this.props.styles.MainContentFadeEnter,
      enterActive: this.props.styles.MainContentFadeEnterActive,
      enterDone: this.props.styles.MainContentFadeEnterDone,
      exit: this.props.styles.MainContentFadeExit,
      exitActive: this.props.styles.MainContentFadeExitActive,
      exitDone: this.props.styles.MainContentFadeExitDone
    };
    const path = this.props.location.pathname;
    return (
      <TransitionGroup styleName="Background">
        <CSSTransition key={path} appear enter exit mountOnEnter unmountOnExit timeout={300} classNames={mainContentFade}
          onEnter={()=>{
            console.info('%s onEnter', path);
          }}
          onEntering={()=>{
            console.info('%s onEntering', path);
          }}
          onEntered={()=>{
            console.info('%s onEntered', path);
          }}
          onExit={()=>{
            console.info('%s onExit', path);
          }}
          onExiting={()=>{
            console.info('%s onExiting', path);
          }}
          onExited={()=>{
            console.info('%s onExited', path);
          }}
        >
          <Switch>
            <Route path="/login" component={Login} />
            <Route exact path="/" component={AppList} />
            <Route path="/app/settings" component={Settings} />
            <Route path="/app/animation" component={Animation} />
            <Route appName="Server" path="/app/server" hostname={window.location.hostname} component={ServerContainer} />
          </Switch>
        </CSSTransition>
      </TransitionGroup>
    );
  }
}

function mapStateToProps(state) {
  return {
    location: state.routing.location
  };
}

Desktop.propTypes = {
  location: PropTypes.object,
  styles: PropTypes.object,
};

export default connect(mapStateToProps)(CSSModules(Desktop, styles));
