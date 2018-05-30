import React from 'react';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import { CSSTransition, TransitionGroup } from 'react-transition-group';
import styles from './Animation.css';

class AppGroup extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const iconStyles = {
      appear: this.props.styles.AppGroupAppear,
      appearActive: this.props.styles.AppGroupAppearActive,
      enter: this.props.styles.AppGroupEnter,
      enterActive: this.props.styles.AppGroupEnterActive,
      enterDone: this.props.styles.AppGroupEnterDone,
      exit: this.props.styles.AppGroupExit,
      exitAction: this.props.styles.AppGroupExitAction,
      exitDone: this.props.styles.AppGroupExitDone
    };
    return (
      <CSSTransition classNames={iconStyles} key={'AppGroup'} appear in exit timeout={500}
        onEnter={()=>{console.info('AppGroup onEnter');}}
        onEntering={()=>{console.info('AppGroup onEntering');}}
        onEntered={()=>{console.info('AppGroup onEntered');}}
        onExit={()=>{console.info('AppGroup onExit');}}
        onExiting={()=>{console.info('AppGroup onExiting');}}
        onExited={()=>{console.info('AppGroup onExited');}}
      >
        <TransitionGroup>
          {this.props.children}
        </TransitionGroup>
      </CSSTransition>
    );
  }
}

AppGroup.propTypes = {
  children: PropTypes.array,
  styles: PropTypes.object,
};

export default CSSModules(AppGroup, styles);
