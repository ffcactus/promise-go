import React from 'react';
import { Link } from 'react-router-dom';
function Server() {
  return (
    <React.Fragment>
      <p>Server</p>
      <Link to="/">Home</Link>
    </React.Fragment>
  );
}

export default Server;

