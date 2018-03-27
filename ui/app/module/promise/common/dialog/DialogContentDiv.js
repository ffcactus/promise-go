import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Dialog.css';

const DialogContentDiv = (props) => {
  return (
    <div styleName="DialogContentDiv">
      {props.children}
    </div>
  );
};

DialogContentDiv.propTypes = {
  children: PropTypes.oneOfType([PropTypes.object, PropTypes.array]),
};

export default CSSModules(DialogContentDiv, styles);
