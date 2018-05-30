import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Animation.css';
import DesktopTest from './DesktopTest';
import AppGroup from './AppGroup';
import IconTest from './IconTest';
import AppContainer from './AppContainer';

class AnimationTest extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      desktop: false,
      app: false
    };
    this.changeContent = this.changeContent.bind(this);
  }

  changeContent(event) {
    event.preventDefault();
    switch (event.currentTarget.value) {
      case 'DesktopIn':
        this.setState((state) => {
          return {
            ...state,
            desktop: true
          };
        });
        break;
      case 'DesktopOut':
        this.setState((state) => {
          return {
            ...state,
            desktop: false
          };
        });
        break;
      case 'AppIn':
        this.setState((state) => {
          return {
            ...state,
            app: true
          };
        });
        break;
      case 'AppOut':
        this.setState((state) => {
          return {
            ...state,
            app: false
          };
        });
        break;
      case 'Switch':
        this.setState((state) => {
          return {
            ...state,
            app: !state.app,
            desktop: !state.desktop
          };
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
    // appear enter exit mountOnEnter unmountOnExit
    return (
      <div>
        <DesktopTest>
          <AppGroup key={'AppGroup'} inProp={this.state.desktop} >
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
          <AppContainer key={'AppContainer'} inProp={this.state.app}/>
        </DesktopTest>
        <input type="button" value="DesktopIn" onClick={this.changeContent}/>
        <input type="button" value="DesktopOut" onClick={this.changeContent}/>
        <br/>
        <input type="button" value="AppIn" onClick={this.changeContent}/>
        <input type="button" value="AppOut" onClick={this.changeContent}/>
        <br/>
        <input type="button" value="Switch" onClick={this.changeContent}/>
      </div>
    );
  }
}

export default CSSModules(AnimationTest, styles);
