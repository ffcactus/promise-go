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
      <div styleName="flex-item flex-row-container border-column" style={{maxHeight: '40px'}}>
        <ServerGroupControlCreate />
        <ServerGroupControlEdit />
      </div>
    );
  }
}

export default CSSModules(ServerGroupControlArea, styles, {allowMultiple: true});

