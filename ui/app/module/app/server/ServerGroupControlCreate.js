import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

class ServerGroupControlCreate extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const icon = require('./img/icon/Create.png');
    return (
      <div styleName="ServerGroupControlButton">
        <img src={icon}/>
      </div>
    );
  }
}

export default CSSModules(ServerGroupControlCreate, styles);

