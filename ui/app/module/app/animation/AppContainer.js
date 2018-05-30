import React from 'react';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import { CSSTransition } from 'react-transition-group';
import styles from './Animation.css';

class AppContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const transitionStyles = {
      appear: this.props.styles.AppContainerAppear,
      appearActive: this.props.styles.AppContainerAppearActive,
      enter: this.props.styles.AppContainerEnter,
      enterActive: this.props.styles.AppContainerEnterActive,
      enterDone: this.props.styles.AppContainerEnterDone,
      exit: this.props.styles.AppContainerExit,
      exitAction: this.props.styles.AppContainerExitAction,
      exitDone: this.props.styles.AppContainerExitDone
    };
    return (
      <CSSTransition classNames={transitionStyles} key={'AppContainer'} appear in exit timeout={500}
        onEnter={()=>{console.info('AppContainer onEnter');}}
        onEntering={()=>{console.info('AppContainer onEntering');}}
        onEntered={()=>{console.info('AppContainer onEntered');}}
        onExit={()=>{console.info('AppContainer onExit');}}
        onExiting={()=>{console.info('AppContainer onExiting');}}
        onExited={()=>{console.info('AppContainer onExited');}}
      >
        <div>
          <div>
            <h1>This is my App.</h1>
            <p>My app mock iOS interface.</p>
          </div>
        </div>
      </CSSTransition>
    );
  }
}

AppContainer.propTypes = {
  children: PropTypes.object,
  styles: PropTypes.object,
};

export default CSSModules(AppContainer, styles);
