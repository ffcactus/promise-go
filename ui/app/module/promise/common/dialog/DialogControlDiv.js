import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Dialog.css';

const DialogControlDiv = (props) => {
  return (
    <div styleName="DialogControlDiv">
      {props.children}
    </div>
  );
};

DialogControlDiv.propTypes = {
  children: PropTypes.oneOfType([PropTypes.object, PropTypes.array]),
};

export default CSSModules(DialogControlDiv, styles);
