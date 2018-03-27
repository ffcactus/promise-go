import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Dialog.css';

const DialogTitle = (props) => {
  return (
    <div styleName="DialogTitleDiv">
      <p styleName="DialogTitleText">{props.value}</p>
    </div>
  );
};

DialogTitle.propTypes = {
  value: PropTypes.string
};

export default CSSModules(DialogTitle, styles);
