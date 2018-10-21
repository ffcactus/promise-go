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
import EnclosureContainer from '../../app/enclosure/EnclosureContainer';
import styles from './Desktop.css';
import ServerController from '../../app/server-controller/ServerController';

class Desktop extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
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
      default:
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
    }
    return (
      <TransitionGroup styleName="Background">
        <CSSTransition key={path} appear enter exit mountOnEnter unmountOnExit timeout={300} classNames={animationStyle}>
          <Switch location={this.props.location}>
            <Route exact path="/login" component={Login} />
            <PrivateRoute exact path="/" component={AppList} />
            <PrivateRoute path="/settings" component={Settings} />
            <PrivateRoute path="/test" component={Test} />
            <PrivateRoute path="/animation" component={Animation} />
            <PrivateRoute path="/servercontroller" component={ServerController} />
            <PrivateRoute appName="Server" path="/server" hostname={window.location.hostname} component={ServerContainer} />
            <PrivateRoute appName="Enclosure" path="/enclosure" hostname={window.location.hostname} component={EnclosureContainer} />
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
