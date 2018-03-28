import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerGroupControlCreate from './ServerGroupControlCreate';
import ServerGroupControlEdit from './ServerGroupControlEdit';

class ServerListControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="ServerListControlArea">
      </div>
    );
  }
}

export default CSSModules(ServerListControlArea, styles);

