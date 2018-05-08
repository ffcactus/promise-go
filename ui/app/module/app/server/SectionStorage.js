import React from 'react';
import PropTypes from 'prop-types';
import SectionHeader from './SectionHeader';
import ObjectStorage from './ObjectStorage';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

function SectionStorage(props) {
  const iconImage = require('./img/icon/Storage.png');
  return (<div styleName="ServerDetailSectionDiv">
    <SectionHeader name="Storage" image={iconImage} />
    {
      props.storages.map((each, i) => {
        return <ObjectStorage key={i.toString()} storage={each} />;
      })
    }
  </div>);
}

SectionStorage.propTypes = {
  storages: PropTypes.array
};

export default CSSModules(SectionStorage, styles);

