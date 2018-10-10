import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Enclosure.css';

class ResourceSearch extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <input styleName="search" placeholder="Search resource"/>
    );
  }
}

export default CSSModules(ResourceSearch, styles, {allowMultiple: true});
