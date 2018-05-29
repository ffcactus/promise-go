import React from 'react';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import { TransitionGroup } from 'react-transition-group';
import styles from './Phone.css';

class DesktopTest extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <TransitionGroup styleName="Desktop">
        {this.props.children}
      </TransitionGroup>
    );
  }
}

DesktopTest.propTypes = {
  children: PropTypes.any,
};

export default CSSModules(DesktopTest, styles);
