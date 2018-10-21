import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './App.css';

class ProfileList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container border-column flex-item-last">
        <p>Profile</p>
      </div>
    );
  }
}

export default CSSModules(ProfileList, styles, {allowMultiple: true});
