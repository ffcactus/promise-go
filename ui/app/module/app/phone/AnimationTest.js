import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Phone.css';

class AnimationTest extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const image1 = require('./img/icon/Phone.png');
    const image2 = require('./img/icon/Phone.png');
    return (
      <div>
        <div styleName="Desktop">
          <div styleName="IconDiv">
            <img src={image1}/>
          </div>
          <div styleName="IconDiv">
            <img src={image2}/>
          </div>
        </div>
        <input type="button" value="App"/>
        <input type="button" value="Desktop"/>
      </div>
    );
  }
}

export default CSSModules(AnimationTest, styles);
