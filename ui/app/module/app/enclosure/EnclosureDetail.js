import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './App.css';

const EnclosureDetail = props => {
  if (!props.enclosure) {
    return <div />;
  }
  return (
    <div styleName="detail-container">
      <h1>Name</h1>
      <p>{props.enclosure.Name}</p>
      <h2>Description</h2>
      <p>{props.enclosure.Description}</p>
    </div>
  );
};

EnclosureDetail.propTypes = {
  enclosure: PropTypes.object,
};

export default connect()(CSSModules(EnclosureDetail, styles, {allowMultiple: true}));
