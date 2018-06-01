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
import Test from '../../app/test/Test';
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
    let animationStyle;
    switch(path) {
      case '/':
        animationStyle = {
          appear: this.props.styles.AppListAppear,
          appearActive: this.props.styles.AppListAppearActive,
          enter: this.props.styles.AppListEnter,
          enterActive: this.props.styles.AppListEnterActive,
          enterDone: this.props.styles.AppListEnterDone,
          exit: this.props.styles.AppListExit,
          exitActive: this.props.styles.AppListExitActive,
          exitDone: this.props.styles.AppListExitDone
        };
        break;
      case '/animation':
        animationStyle = {
          appear: this.props.styles.MainContentFadeAppear,
          appearActive: this.props.styles.MainContentFadeAppearActive,
          enter: this.props.styles.MainContentFadeEnter,
          enterActive: this.props.styles.MainContentFadeEnterActive,
          enterDone: this.props.styles.MainContentFadeEnterDone,
          exit: this.props.styles.MainContentFadeExit,
          exitActive: this.props.styles.MainContentFadeExitActive,
          exitDone: this.props.styles.MainContentFadeExitDone
        };
        break;
      case '/test':
        animationStyle = {
          appear: this.props.styles.TestAppear,
          appearActive: this.props.styles.TestAppearActive,
          enter: this.props.styles.TestEnter,
          enterActive: this.props.styles.TestEnterActive,
          enterDone: this.props.styles.TestEnterDone,
          exit: this.props.styles.TestExit,
          exitActive: this.props.styles.TestExitActive,
          exitDone: this.props.styles.TestExitDone
        };
        break;
      default:
        break;
    }
    return (
      <TransitionGroup styleName="Background">
        <CSSTransition key={path} appear enter exit mountOnEnter unmountOnExit timeout={300} classNames={animationStyle}
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
            <Route path="/settings" component={Settings} />
            <Route path="/test" component={Test} />
            <Route path="/animation" component={Animation} />
            <Route appName="Server" path="/server" hostname={window.location.hostname} component={ServerContainer} />
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
