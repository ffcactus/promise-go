import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerSortOrderSelect from './ServerSortOrderSelect';

class ServerListControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="ServerListControlArea" >
        <ServerSortOrderSelect options={['Name', 'Health']} />
      </div>
    );
  }
}

export default CSSModules(ServerListControlArea, styles);
