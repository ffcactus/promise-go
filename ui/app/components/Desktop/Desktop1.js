import React from 'react';
import PropTypes from 'prop-types';
import AppIcon from './AppIcon';
import CSSModules from 'react-css-modules';
import styles from '../../styles/Desktop/Desktop.css';


function Desktop() {
  const appIcons = [
    {
      'name': 'Settings',
      'image': require('../../img/icon/Settings.png'),
      'notificationCount': 0
    },
    {
      'name': 'Phone',
      'image': require('../../img/icon/Phone.png'),
      'notificationCount': 999
    },
    {
      'name': 'Photos',
      'image': require('../../img/icon/Photos.png'),
      'notificationCount': 0
    },
    {
      'name': 'Podcast',
      'image': require('../../img/icon/Podcast.png'),
      'notificationCount': 10000
    },
    {
      'name': 'Maps',
      'image': require('../../img/icon/Maps.png'),
      'notificationCount': 0
    },
    {
      'name': 'Clock',
      'image': require('../../img/icon/Clock.png'),
      'notificationCount': 0
    },
    {
      'name': 'Mail',
      'image': require('../../img/icon/Mail.png'),
      'notificationCount': 7
    }
  ];

  return (
    <div styleName="desktop">
      <div styleName="appArea">
        {appIcons.map(each => {
          return <AppIcon key={each.name} name={each.name} image={each.image} notificationCount={each.notificationCount} />;
        })}
      </div>
    </div>
  );
}

Desktop.propTypes = {
  appIcons: PropTypes.object
};

export default CSSModules(Desktop, styles);
