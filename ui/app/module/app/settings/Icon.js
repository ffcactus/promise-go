import React from 'react';
import AppIcon from '../../promise/common/AppIcon';
// import PropTypes from 'prop-types';

function Icon() {
  const name = 'Settings';
  const image = require('./img/icon/Settings.png');
  const notificationCount = 0;
  const uri = '/app/settings';
  return <AppIcon key={name} name={name} uri={uri} image={image} notificationCount={notificationCount} />;
}

export default Icon;
