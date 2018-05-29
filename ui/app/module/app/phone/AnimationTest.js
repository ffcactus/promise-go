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
            <IconTest img={image1} name="Server"/>
            <IconTest img={image1} name="Server"/>
            <IconTest img={image1} name="Enclosure"/>
            <IconTest img={image1} name="Pool"/>
            <IconTest img={image1} name="Task"/>
            <IconTest img={image1} name="Settings"/>
          </AppGroup>
        </DesktopTest>
        <input type="button" value="App"/>
        <input type="button" value="Desktop"/>
      </div>
    );
  }
}

export default CSSModules(AnimationTest, styles);
