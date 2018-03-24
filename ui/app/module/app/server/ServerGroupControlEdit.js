import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';

class ServerGroupControlEdit extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const icon = require('./img/icon/Edit.png');
    return (
      <div styleName="ServerGroupControlButton">
        <img src={icon}/>
      </div>
    );
  }
}

export default CSSModules(ServerGroupControlEdit, styles);

