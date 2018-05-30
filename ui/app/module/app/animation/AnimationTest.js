import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Animation.css';
import DesktopTest from './DesktopTest';
import AppGroup from './AppGroup';
import IconTest from './IconTest';
import AppContainer from './AppContainer';
import { CSSTransition } from 'react-transition-group';

class AnimationTest extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      content: 'AppGroup'
    };
    this.changeContent = this.changeContent.bind(this);
  }

  changeContent(event) {
    event.preventDefault();
    switch (event.currentTarget.value) {
      case 'AppGroup':
        this.setState({
          content: 'AppGroup'
        });
        break;
      case 'AppContainer':
        this.setState({
          content: 'AppContainer'
        });
        break;
      default:
        break;
    }
  }

  render() {
    const Clock = require('./img/icon/Clock.png');
    const Mail = require('./img/icon/Mail.png');
    const Maps = require('./img/icon/Maps.png');
    const Photos = require('./img/icon/Photos.png');
    const Podcast = require('./img/icon/Podcast.png');
    const Phone = require('./img/icon/Phone.png');
    const Settings = require('./img/icon/Settings.png');

    let current = null;

    switch(this.state.content) {
      case 'AppGroup':
        current = (
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
        );
        break;
      case 'AppContainer':
        current = (
          <AppContainer/>
        );
        break;
      default:
        break;
    }

    return (
      <div>
        <DesktopTest>
          {current}
        </DesktopTest>
        <input type="button" value="AppGroup" onClick={this.changeContent}/>
        <input type="button" value="AppContainer" onClick={this.changeContent}/>
      </div>
    );
  }
}

export default CSSModules(AnimationTest, styles);
