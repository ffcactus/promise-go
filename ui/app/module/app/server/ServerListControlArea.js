import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

class ServerListControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="ServerListControlArea" />
    );
  }
}

export default CSSModules(ServerListControlArea, styles);

