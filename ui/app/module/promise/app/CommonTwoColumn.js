import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';

function CommonTwoColumn(props) {
  return (
    <div styleName="CommonTwoColumn">
      <div styleName="CommonTwoColumnLeft">
        {props.left}
      </div>
      <div styleName="CommonTwoColumnRight">
        {props.right}
      </div>
    </div>
  );
}

CommonTwoColumn.propTypes = {
  left: PropTypes.object,
  right: PropTypes.object
};

export default CSSModules(CommonTwoColumn, Styles);
