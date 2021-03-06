import React from 'react';
import PropTypes from 'prop-types';

export const Health = (props) => {
  let color = 'gray';
  switch (props.health) {
    case 'OK':
      color = 'limegreen';
      break;
    case 'Warning':
      color = 'orange';
      break;
    case 'Critical':
      color = 'red';
      break;
    default:
      color = 'gray';
      break;
  }
  return (
    <div style={{
      margin: '20px',
      height: '10px',
      width: '10px',
      backgroundColor: color,
      borderRadius: '10px',
    }}/>
  );
};

Health.propTypes = {
  health: PropTypes.string,
};
