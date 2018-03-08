import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../styles/ServerFrame.css';

function Startup(props) {
  return (<div>
    <table>
      <thead>
        <tr>
          <th styleName="level1">Process</th>
          <th styleName="level1">State</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>Websocket</td>
          <td>{props.startup.ws.state}</td>
        </tr>
      </tbody>
    </table>
  </div>);
}

Startup.propTypes = {
  startup: PropTypes.object
};

export default CSSModules(Startup, styles);
