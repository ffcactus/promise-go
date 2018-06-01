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
      <TransitionGroup>
        <AppIcon img={Clock} uri="/test" name="Test" notificationCount={0}/>
        <AppIcon img={Mail} uri="/settings" name="Mail" notificationCount={1}/>
        <AppIcon img={Maps} uri="/settings" name="Maps" notificationCount={0}/>
        <AppIcon img={Photos} uri="/settings" name="Photos" notificationCount={10000}/>
        <AppIcon img={Podcast} uri="/settings" name="Podcast" notificationCount={0}/>
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
