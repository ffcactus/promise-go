import React from 'react';
import PropTypes from 'prop-types';
import ServerDetailSectionHeader from './ServerDetailSectionHeader';
import ServerDetailObjectStorage from './ServerDetailObjectStorage';
import CSSModules from 'react-css-modules';
import styles from './App.css';

function ServerDetailSectionStorage(props) {
  const iconImage = require('./img/icon/Storage.png');
  return (<div styleName="ServerDetailSectionDiv">
    <ServerDetailSectionHeader name="Storage" image={iconImage} />
    {
      props.storages.map((each, i) => {
        return <ServerDetailObjectStorage key={i.toString()} storage={each} />;
      })
    }
  </div>);
}

ServerDetailSectionStorage.propTypes = {
  storages: PropTypes.array
};

export default CSSModules(ServerDetailSectionStorage, styles);

