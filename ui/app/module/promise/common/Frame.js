import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/frame.css';
import Menu from './Menu';

function Frame(props) {
  return (
    <div>
      <div styleName="header"><Menu /></div>
      <div styleName="main">{props.main}</div>
      <div styleName="footer">{props.footer}</div>
    </div>
  );
}

Frame.propTypes = {
  main: PropTypes.object,
  footer: PropTypes.object
};

export default CSSModules(Frame, styles);
