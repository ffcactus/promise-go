import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';

function SectionBasic(props) {
  return (<div styleName="sectionDiv">
    <h5>Name</h5>
    <p>{props.server.Name}</p>
    <h5>Description</h5>
    <p>{props.server.Description}</p>
    <h5>Type</h5>
    <p>{props.server.Type}</p>
    <h5>State</h5>
    <p>{props.server.State}</p>
    <h5>Health</h5>
    <p>{props.server.Health}</p>
  </div>);
}

SectionBasic.propTypes = {
  server: PropTypes.object
};

export default CSSModules(SectionBasic, styles);
