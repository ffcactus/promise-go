import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerSearchArea from './ServerSearchArea';
import ServerListControlArea from './ServerListControlArea';
import ServerList from './ServerListNew';

class ServerListArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="ServerListArea">
        <ServerSearchArea />
        <ServerListControlArea />
        <div style={{ flex: '1 1 auto' }}>
          <ServerList />
        </div>
      </div>
    );
  }
}

export default CSSModules(ServerListArea, styles);

