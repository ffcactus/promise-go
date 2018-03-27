import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Dialog.css';

const DialogButton = (props) => {
  return (
    <button onClick={props.onClick} styleName="DialogButton">{props.name}</button>
  );
};

DialogButton.propTypes = {
  name: PropTypes.string,
  onClick: PropTypes.func
};

export default CSSModules(DialogButton, styles);
