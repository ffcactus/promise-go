import React from 'react';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import { CSSTransition } from 'react-transition-group';
import styles from './Phone.css';

class MyApp extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const transitionStyles = {
      appear: this.props.styles.MyAppAppear,
      appearActive: this.props.styles.MyAppAppearActive,
      enter: this.props.styles.MyAppEnter,
      enterActive: this.props.styles.MyAppEnterActive,
      enterDone: this.props.styles.MyAppEnterDone,
    };
    return (
      <CSSTransition classNames={transitionStyles} in appear timeout={500}>
        <div>
          <h1>This is my App.</h1>
          <p>My app mock iOS interface.</p>
        </div>
      </CSSTransition>
    );
  }
}

MyApp.propTypes = {
  children: PropTypes.object,
  styles: PropTypes.object,
};

export default CSSModules(MyApp, styles);
