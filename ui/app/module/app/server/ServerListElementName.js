import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

function ServerListElementName(props) {
  return (
    <div styleName="ServerListElementName">
      <p>{props.name}</p>
    </div>
  );
}

ServerListElementName.propTypes = {
  name: PropTypes.string,
};

export default CSSModules(ServerListElementName, styles);
