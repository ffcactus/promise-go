import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Enclosure.css';

function EnclosureDetailArea() {
  return (
    <div styleName="flex-column-container detail-area border-row"/>
  );
}

export default CSSModules(EnclosureDetailArea, styles, {allowMultiple: true});
