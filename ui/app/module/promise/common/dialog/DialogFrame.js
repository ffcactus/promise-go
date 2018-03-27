import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Dialog.css';

const DialogFrame = (props) => {
  return (
    <div styleName="DialogFrameContainer">
      <div styleName="DialogFrame">
        {props.children}
      </div>
    </div>
  );
};


DialogFrame.propTypes = {
  children: PropTypes.oneOfType([PropTypes.object, PropTypes.array]),
};

export default CSSModules(DialogFrame, styles);
