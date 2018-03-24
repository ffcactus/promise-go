import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerGroupControlCreate from './ServerGroupControlCreate';
import ServerGroupControlEdit from './ServerGroupControlEdit';

class ServerGroupControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="ServerGroupControlArea">
        <ServerGroupControlCreate />
        <ServerGroupControlEdit />
      </div>
    );
  }
}

export default CSSModules(ServerGroupControlArea, styles);

