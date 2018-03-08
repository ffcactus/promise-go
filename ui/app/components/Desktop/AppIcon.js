import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import IconNotification from './IconNotification';
import styles from '../../styles/Desktop/Desktop.css';

function AppIcon(props) {
  return (
    <div styleName="AppIcon">
      <div styleName="AppIconAndNotification">
        <img src={props.image} />
        <IconNotification notificationCount={props.notificationCount} />
      </div>
      <div styleName="AppIconName">
        <p>{props.name}</p>
      </div>
    </div>
  );
}

AppIcon.propTypes = {
  name: PropTypes.string,
  image: PropTypes.string,
  notificationCount: PropTypes.number
};

export default CSSModules(AppIcon, styles);
