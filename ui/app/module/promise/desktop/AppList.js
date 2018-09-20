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
    const Contacts = require('./img/icon/contacts_icon@2x.png');
    const Date = require('./img/icon/date_icon@2x.png');
    const Find = require('./img/icon/find_icon@2x.png');
    const Friends = require('./img/icon/fmf_icon@2x.png');
    const Drive = require('./img/icon/icloud_drive_icon@2x.png');
    const Keynote = require('./img/icon/keynote_icon@2x.png');
    const Mail = require('./img/icon/mail_icon@2x.png');
    const Notes = require('./img/icon/notes_icon@2x.png');
    const Numbers = require('./img/icon/numbers_icon@2x.png');
    const Pages = require('./img/icon/pages_icon@2x.png');
    const Photos = require('./img/icon/photos_icon@2x.png');
    const Reminders = require('./img/icon/reminders_icon@2x.png');
    const Settings = require('./img/icon/settings_icon@2x.png');

    return (
      <TransitionGroup styleName="AppList" key="AppList" _id="AppList">
        <AppIcon img={Contacts} uri="/test" name="Contacts" notificationCount={0}/>
        <AppIcon img={Date} uri="/servercontroller" name="Date" notificationCount={0}/>
        <AppIcon img={Find} uri="/login" name="Find" notificationCount={0}/>
        <AppIcon img={Friends} uri="/animation" name="Friends" notificationCount={0}/>
        <AppIcon img={Drive} uri="/server" name="Server" notificationCount={0}/>
        <AppIcon img={Keynote} uri="/server" name="Keynote" notificationCount={0}/>
        <AppIcon img={Mail} uri="/settings" name="Mail" notificationCount={0}/>
        <AppIcon img={Notes} uri="/settings" name="Notes" notificationCount={0}/>
        <AppIcon img={Numbers} uri="/settings" name="Numbers" notificationCount={0}/>
        <AppIcon img={Pages} uri="/settings" name="Pages" notificationCount={0}/>
        <AppIcon img={Photos} uri="/settings" name="Photos" notificationCount={0}/>
        <AppIcon img={Reminders} uri="/settings" name="Reminders" notificationCount={0}/>
        <AppIcon img={Settings} uri="/settings" name="Settings" notificationCount={0}/>
      </TransitionGroup>
    );
  }
}

AppList.propTypes = {
  styles: PropTypes.object,
};

export default CSSModules(AppList, styles);
