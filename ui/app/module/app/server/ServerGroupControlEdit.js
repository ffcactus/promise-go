import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import CenterDiv from '../../promise/common/CenterDiv';

class ServerGroupControlEdit extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const icon = require('../../promise/common/img/icon/Navigation_Compose_2x.png');
    return (
      <div styleName="ServerGroupControlButton">
        <CenterDiv><img src={icon} style={{display: 'block', margin: 'auto', height: '30px'}}/></CenterDiv>
      </div>
    );
  }
}

export default CSSModules(ServerGroupControlEdit, styles);

