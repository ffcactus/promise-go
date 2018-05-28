import React from 'react';
import CSSModules from 'react-css-modules';
import { CSSTransition } from 'react-transition-group';
import styles from './Phone.css';

class AnimationTest extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const image1 = require('./img/icon/Phone.png');
    const iconStyles = {
      appear: this.props.styles.IconAppear,
      appearActive: this.props.styles.IconAppearActive,
      enter: this.props.styles.IconEnter,
      enterActive: this.props.styles.IconEnterActive,
      enterDone: this.props.styles.IconEnterDone,
    };
    const appsStyles = {
      appear: this.props.styles.AppsAppear,
      appearActive: this.props.styles.AppsAppearActive,
      enter: this.props.styles.AppsEnter,
      enterActive: this.props.styles.AppsEnterActive,
      enterDone: this.props.styles.AppsEnterDone,
    };
    const apps =
      <CSSTransition classNames={appsStyles} in appear timeout={500}>
        <div>
          <CSSTransition classNames={iconStyles} in appear timeout={500}>
            <div>
              <img src={image1}/>
            </div>
          </CSSTransition>
          <CSSTransition classNames={iconStyles} in appear timeout={500}>
            <div>
              <img src={image1}/>
            </div>
          </CSSTransition>
          <CSSTransition classNames={iconStyles} in appear timeout={500}>
            <div>
              <img src={image1}/>
            </div>
          </CSSTransition>
          <CSSTransition classNames={iconStyles} in appear timeout={500}>
            <div>
              <img src={image1}/>
            </div>
          </CSSTransition>
          <CSSTransition classNames={iconStyles} in appear timeout={500}>
            <div>
              <img src={image1}/>
            </div>
          </CSSTransition>
        </div>
      </CSSTransition>;

    return (
      <div>
        <div styleName="Desktop">
          {apps}
        </div>
        <input type="button" value="App"/>
        <input type="button" value="Desktop"/>
      </div>
    );
  }
}

export default CSSModules(AnimationTest, styles);
