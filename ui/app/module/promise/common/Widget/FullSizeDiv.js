import React from 'react';
import PropTypes from 'prop-types';

const _style = {
  height: '100%',
  width: '100%',
  position: 'absolute'
};

const FullSizeDiv = (props) => (<div style={_style}>{props.children}</div>);

FullSizeDiv.propTypes = {
  children: PropTypes.any,
};

export default FullSizeDiv;
