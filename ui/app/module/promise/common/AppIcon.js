import React from 'react';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import IconNotification from './IconNotification';
import styles from './AppIcon.css';

function AppIcon(props) {
  return (
    <div styleName="AppIcon">
      <div styleName="AppIconAndNotification">
        <Link to={props.uri}>
          <img src={props.image} />
          <IconNotification notificationCount={props.notificationCount} />
        </Link>
      </div>
      <div styleName="AppIconName">
        <p>{props.name}</p>
      </div>
    </div>
  );
}

AppIcon.propTypes = {
  name: PropTypes.string,
  uri: PropTypes.string,
  image: PropTypes.string,
  notificationCount: PropTypes.number
};

export default CSSModules(AppIcon, styles);
