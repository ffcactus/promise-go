import React from 'react';
import { Link } from 'react-router-dom';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import styles from './Test.css';

class Test extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <p>Settings</p>
        <Link to="/">Home</Link>
        <input type="button" value="Switch"/>
      </div>
    );
  }
}

Test.propTypes = {
  styles: PropTypes.object,
};

export default CSSModules(Test, styles);

