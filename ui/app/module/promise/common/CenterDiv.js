import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './CenterDiv.css';

const CenterDiv = (props) => {
  return (
    <div styleName="WarpOuter">
      <div styleName="WarpInner">{props.children}</div>
    </div>
  );
};

CenterDiv.propTypes = {
  children: PropTypes.object,
};

export default CSSModules(CenterDiv, styles);
