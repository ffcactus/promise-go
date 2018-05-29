import React from 'react';
import { Link } from 'react-router-dom';
import AnimationTest from './AnimationTest';
function Animation() {
  return (
    <React.Fragment>
      <p>Animation</p>
      <Link to="/">Home</Link>
      <AnimationTest />
    </React.Fragment>
  );
}

export default Animation;
