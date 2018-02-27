import React from 'react';
import PropTypes from 'prop-types';
import SettingsAppIcon from '../../app/settings/Icon';
import PhoneAppIcon from '../../app/phone/Icon';
import Server from '../../app/server/Icon';
import CSSModules from 'react-css-modules';
import styles from './Desktop.css';


function AppCollection() {
  return (
    <div styleName="desktop">
      <div styleName="appArea">
        <SettingsAppIcon />
        <PhoneAppIcon />
        <Server />
      </div>
    </div>
  );
}

AppCollection.propTypes = {
  appIcons: PropTypes.string
};

export default CSSModules(AppCollection, styles);
