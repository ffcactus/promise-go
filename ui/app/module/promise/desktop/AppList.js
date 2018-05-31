import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import { TransitionGroup, CSSTransition } from 'react-transition-group';
import styles from './Desktop.css';
import FullSizeDiv from '../common/Widget/FullSizeDiv';
import AppIcon from './AppIcon';

class AppList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const appListStyle = {
      appear: this.props.styles.AppListAppear,
      appearActive: this.props.styles.AppListAppearActive,
      enter: this.props.styles.AppListEnter,
      enterActive: this.props.styles.AppListEnterActive,
      enterDone: this.props.styles.AppListEnterDone,
      exit: this.props.styles.AppListExit,
      exitActive: this.props.styles.AppListExitActive,
      exitDone: this.props.styles.AppListExitDone
    };

    const Clock = require('./img/icon/Clock.png');
    const Mail = require('./img/icon/Mail.png');
    const Maps = require('./img/icon/Maps.png');
    const Photos = require('./img/icon/Photos.png');
    const Podcast = require('./img/icon/Podcast.png');
    const Phone = require('./img/icon/Phone.png');
    const Settings = require('./img/icon/Settings.png');
    return (
      <div> {/* this div will be used by CSSTransition container. */}
        <FullSizeDiv>
          <CSSTransition appear enter exit mountOnEnter unmountOnExit timeout={300} classNames={appListStyle}
            onEnter={()=>{
              console.info('%s onEnter', 'AppList');
            }}
            onEntering={()=>{
              console.info('%s onEntering', 'AppList');
            }}
            onEntered={()=>{
              console.info('%s onEntered', 'AppList');
            }}
            onExit={()=>{
              console.info('%s onExit', 'AppList');
            }}
            onExiting={()=>{
              console.info('%s onExiting', 'AppList');
            }}
            onExited={()=>{
              console.info('%s onExited', 'AppList');
            }}
          >
            <TransitionGroup>
              <AppIcon img={Clock} name="Clock" notificationCount={0}/>
              <AppIcon img={Mail} name="Mail" notificationCount={1}/>
              <AppIcon img={Maps} name="Maps" notificationCount={0}/>
              <AppIcon img={Photos} name="Photos" notificationCount={10000}/>
              <AppIcon img={Podcast} name="Podcast" notificationCount={0}/>
              <AppIcon img={Phone} name="Phone" notificationCount={0}/>
              <AppIcon img={Settings} name="Settings" notificationCount={0}/>
              <AppIcon img={Photos} name="Photos" notificationCount={2}/>
              <AppIcon img={Maps} name="Maps" notificationCount={0}/>
              <AppIcon img={Mail} name="Mail" notificationCount={1}/>
              <AppIcon img={Podcast} name="Podcast" notificationCount={0}/>
            </TransitionGroup>
          </CSSTransition>
        </FullSizeDiv>
      </div>
    );
  }
}

AppList.propTypes = {
  styles: PropTypes.object,
};

export default CSSModules(AppList, styles);
