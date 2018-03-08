import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/dialog.css';
import CenterDiv from './CenterDiv';

const ConfirmDialog = (props) => {
  return (
    <CenterDiv>
      <div styleName="ConfirmDialog">
        <div styleName="ConfirmTitleDiv">
          <p styleName="ConfirmTitleText">{props.title}</p>
        </div>
        <hr styleName="ConfirmHr" />
        <div styleName="ConfirmContentDiv">
          <p>{props.message}</p>
        </div>
        <div styleName="ConfirmButtonDiv">
          <p styleName="ConfirmButtonText">Confirm</p>
        </div>
      </div>
    </CenterDiv>
  );
};

ConfirmDialog.propTypes = {
  title: PropTypes.string,
  message: PropTypes.string,
  onConfirm: PropTypes.func,
};

ConfirmDialog.propTypes = {
  title: PropTypes.string,
  message: PropTypes.string,
  onConfirm: PropTypes.func,
};

export default CSSModules(ConfirmDialog, styles);
