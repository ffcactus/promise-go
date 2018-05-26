import React from 'react';
import { Link } from 'react-router-dom';
import AnimationTest from './AnimationTest';
function Phone() {
  return (
    <React.Fragment>
      <p>Phone</p>
      <Link to="/">Home</Link>
      <AnimationTest />
    </React.Fragment>
  );
}

export default Phone;
