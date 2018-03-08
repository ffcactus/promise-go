import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';

const SectionHeader = props =>
  <div styleName="detailSectionHeader">
    <img styleName="detailSectionHeaderIcon" src={props.image} />
    <p>{props.name}</p>
  </div>;

SectionHeader.propTypes = {
  image: PropTypes.string,
  name: PropTypes.string
};

export default CSSModules(SectionHeader, styles);
