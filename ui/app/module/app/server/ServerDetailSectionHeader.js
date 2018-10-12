import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './App.css';

const ServerDetailSectionHeader = props =>
  <div styleName="ServerDetailSectionHeader">
    <img styleName="ServerDetailSectionHeaderIcon" src={props.image} />
    <p>{props.name}</p>
  </div>;

ServerDetailSectionHeader.propTypes = {
  image: PropTypes.string,
  name: PropTypes.string
};

export default CSSModules(ServerDetailSectionHeader, styles);
