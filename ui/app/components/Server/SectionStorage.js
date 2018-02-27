import React from 'react';
import PropTypes from 'prop-types';
import SectionHeader from './SectionHeader';
import ObjectStorage from './ObjectStorage';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';

function SectionStorage(props) {
  const iconImage = require('../../img/icon/Storage.png');
  return (<div styleName="sectionDiv">
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

