import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './App.css';

class ServerGroupControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container border-column" style={{maxHeight: '40px'}}/>
    );
  }
}

export default CSSModules(ServerGroupControlArea, styles, {allowMultiple: true});
