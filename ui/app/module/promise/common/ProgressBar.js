import React from 'react';

const ProgressBar = () => {
  return (
    <svg>
      <rect x="0" y="0" width="100" height="14" style={{
        fill: 'white',
        stroke: 'green',
        strokeWidth: '1'
      }} />
      <rect x="0" y="0" width="50" height="14" style={{
        fill: 'green',
        stroke: 'green',
        strokeWidth: '1'
      }} />
    </svg>
  );
};

export { ProgressBar };
