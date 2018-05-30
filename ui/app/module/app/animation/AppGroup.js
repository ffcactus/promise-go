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
      exitActive: this.props.styles.AppGroupExitActive,
      exitDone: this.props.styles.AppGroupExitDone
    };
    return (
      <CSSTransition in={this.props.inProp} appear enter exit mountOnEnter unmountOnExit timeout={300} classNames={iconStyles}
        onEnter={()=>{
          console.info('AppGroup onEnter');
        }}
        onEntering={()=>{
          console.info('AppGroup onEntering');
        }}
        onEntered={()=>{
          console.info('AppGroup onEntered');
        }}
        onExit={()=>{
          console.info('AppGroup onExit');
        }}
        onExiting={()=>{
          console.info('AppGroup onExiting');
        }}
        onExited={()=>{
          console.info('AppGroup onExited');
        }}
      >
        <TransitionGroup>
          {this.props.children}
        </TransitionGroup>
      </CSSTransition>
    );
  }
}

AppGroup.propTypes = {
  inProp: PropTypes.bool,
  children: PropTypes.array,
  styles: PropTypes.object,
};

export default CSSModules(AppGroup, styles);
