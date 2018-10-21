import React from 'react';
import CSSModules from 'react-css-modules';
import ResourceDetailControlArea from './ResourceDetailControlArea';
import EnclosureDetailContainer from './EnclosureDetailContainer';
import styles from './App.css';

function EnclosureDetailArea() {
  return (
    <div styleName="flex-column-container detail-area border-row">
      <ResourceDetailControlArea />
      <EnclosureDetailContainer />
    </div>
  );
}

export default CSSModules(EnclosureDetailArea, styles, {allowMultiple: true});
