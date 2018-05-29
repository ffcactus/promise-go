import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Animation.css';
import DesktopTest from './DesktopTest';
import AppGroup from './AppGroup';
import IconTest from './IconTest';

class AnimationTest extends React.Component {
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
      <div>
        <DesktopTest>
          <AppGroup>
            <IconTest img={Clock} name="Clock" notificationCount={0}/>
            <IconTest img={Mail} name="Mail" notificationCount={1}/>
            <IconTest img={Maps} name="Maps" notificationCount={0}/>
            <IconTest img={Photos} name="Photos" notificationCount={10000}/>
            <IconTest img={Podcast} name="Podcast" notificationCount={0}/>
            <IconTest img={Phone} name="Phone" notificationCount={0}/>
            <IconTest img={Settings} name="Settings" notificationCount={0}/>
            <IconTest img={Photos} name="Photos" notificationCount={2}/>
            <IconTest img={Maps} name="Maps" notificationCount={0}/>
            <IconTest img={Mail} name="Mail" notificationCount={1}/>
            <IconTest img={Podcast} name="Podcast" notificationCount={0}/>
          </AppGroup>
        </DesktopTest>
        <input type="button" value="App"/>
        <input type="button" value="Desktop"/>
      </div>
    );
  }
}

export default CSSModules(AnimationTest, styles);
