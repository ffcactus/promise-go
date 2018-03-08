import React from 'react';
import PropTypes from 'prop-types';

function ServerGroupElement(props) {
  return (
    <div>
      <p>{props.name}</p>
    </div>
  );
}

ServerGroupElement.propTypes = {
  name: PropTypes.string,
};

export default ServerGroupElement;

