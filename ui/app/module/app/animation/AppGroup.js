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
    };
    return (
      <CSSTransition classNames={iconStyles} in appear timeout={500}>
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
