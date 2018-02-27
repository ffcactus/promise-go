import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './AppIcon.css';

function IconNotification(props) {
  if (props.notificationCount === 0) {
    return (<div />);
  }
  let v = props.notificationCount.toString();
  if (props.notificationCount > 999) {
    v = '999+';
  }
  return (
    <div styleName="IconNotification">
      <p styleName="NotificationText">{v}</p>
    </div>
  );
}

IconNotification.propTypes = {
  notificationCount: PropTypes.number
};

export default CSSModules(IconNotification, styles);
