import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import { TransitionGroup } from 'react-transition-group';
import styles from './Desktop.css';
import AppIcon from './AppIcon';

class AppList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const Clock = require('./img/icon/Clock.png');
    const Mail = require('./img/icon/Mail.png');
    const Maps = require('./img/icon/Maps.png');
    const Photos = require('./img/icon/Photos.png');
    const Podcast = require('./img/icon/Podcast.png');
    const Phone = require('./img/icon/Phone.png');
    const Settings = require('./img/icon/Settings.png');
    return (
      <TransitionGroup key="AppList" _id="AppList">
        <AppIcon img={Clock} uri="/test" name="Test" notificationCount={0}/>
        <AppIcon img={Mail} uri="/servercontroller" name="ServerController" notificationCount={1}/>
        <AppIcon img={Maps} uri="/login" name="Login" notificationCount={0}/>
        <AppIcon img={Photos} uri="/animation" name="Animation" notificationCount={10000}/>
        <AppIcon img={Podcast} uri="/server" name="Server" notificationCount={0}/>
        <AppIcon img={Phone} uri="/settings" name="Phone" notificationCount={0}/>
        <AppIcon img={Settings} uri="/settings" name="Settings" notificationCount={0}/>
        <AppIcon img={Photos} uri="/settings" name="Photos" notificationCount={2}/>
        <AppIcon img={Maps} uri="/settings" name="Maps" notificationCount={0}/>
        <AppIcon img={Mail} uri="/settings" name="Mail" notificationCount={1}/>
        <AppIcon img={Podcast} uri="/settings" name="Podcast" notificationCount={0}/>
      </TransitionGroup>
    );
  }
}

AppList.propTypes = {
  styles: PropTypes.object,
};

export default CSSModules(AppList, styles);
