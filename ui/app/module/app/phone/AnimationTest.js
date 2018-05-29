import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './Phone.css';
import DesktopTest from './DesktopTest';
import AppGroup from './AppGroup';
import IconTest from './IconTest';

class AnimationTest extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const image1 = require('./img/icon/Phone.png');
    return (
      <div>
        <DesktopTest>
          <AppGroup>
            <IconTest key="Phone1" img={image1}/>
            <IconTest key="Phone2" img={image1}/>
            <IconTest key="Phone3" img={image1}/>
            <IconTest key="Phone4" img={image1}/>
            <IconTest key="Phone5" img={image1}/>
            <IconTest key="Phone6" img={image1}/>
          </AppGroup>
        </DesktopTest>
        <input type="button" value="App"/>
        <input type="button" value="Desktop"/>
      </div>
    );
  }
}

export default CSSModules(AnimationTest, styles);
