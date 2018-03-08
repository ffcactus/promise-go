import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/dialog.css';

const Dialog = (props) => {
  return (
    <div styleName="popup">
      <h1>{props.title}</h1>
      <p>{props.children}</p>
      <button onClick={props.onCancel}>Cancel</button >
      <button onClick={props.onOk}>OK</button >
    </div>
  );
};

Dialog.propTypes = {
  title: PropTypes.string,
  children: PropTypes.object,
  onCancel: PropTypes.func,
  onOk: PropTypes.func
};

export default CSSModules(Dialog, styles);
