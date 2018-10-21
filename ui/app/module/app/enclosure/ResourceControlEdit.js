import React from 'react';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import CenterDiv from '../../promise/common/CenterDiv';
import styles from './App.css';

/**
 * ResourceControlEdit is the button to edit the resource.
 */
class ResourceControlEdit extends React.Component {
  constructor(props) {
    super(props);
    this.onClick = this.onClick.bind(this);
  }

  onClick(event) {
    event.preventDefault();
  }

  render() {
    const icon = require('../../promise/common/img/icon/Navigation_Compose_2x.png');
    return (
      <div styleName="main-control-button-container" style={{float: 'right'}}>
        <CenterDiv><img src={icon} onClick={this.onClick} style={{display: 'block', margin: 'auto', height: '30px'}}/></CenterDiv>
      </div>
    );
  }
}

export default connect()(CSSModules(ResourceControlEdit, styles));
