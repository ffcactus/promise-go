import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../styles/App.css';

class App extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return <div styleName="App">{this.props.children}</div>;
  }
}

App.propTypes = {
  children: PropTypes.object
};

export default CSSModules(App, styles);
