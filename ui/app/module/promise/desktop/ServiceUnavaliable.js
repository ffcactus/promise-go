import React from 'react';
import CSSModules from 'react-css-modules';
import CenterDiv from '../common/CenterDiv';
import styles from './Desktop.css';

const ServiceUnavaliable = () => {
  return (
    <CenterDiv><p styleName="ServiceUnavaliable">Service Unavaliable</p></CenterDiv>
  );
};

export default CSSModules(ServiceUnavaliable, styles);
