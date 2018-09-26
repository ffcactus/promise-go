import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerControlDiscover from './ServerControlDiscover';
import ServerControlEdit from './ServerControlEdit';

class ServerDetailControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container border-column-first" style={{maxHeight: '40px'}}>
        <ServerControlDiscover />
        <ServerControlEdit />
      </div>
    );
  }
}

export default CSSModules(ServerDetailControlArea, styles, {allowMultiple: true});
