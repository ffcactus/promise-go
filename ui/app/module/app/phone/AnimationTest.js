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
    const icon =
      <CSSTransition
        classNames={{
          appear: this.props.styles.CSSTransitionAppear,
          appearActive: this.props.styles.CSSTransitionAppearActive,
          enter: this.props.styles.CSSTransitionEnter,
          enterActive: this.props.styles.CSSTransitionEnterActive,
          enterDone: this.props.styles.CSSTransitionEnterDone,
        }}
        in
        appear
        timeout={800}
        onEnter={() => {
          console.info('onEnter');
        }}
        onEntering={() => {
          console.info('onEntering');
        }}
        onEntered={() => {
          console.info('onEntered');
        }}>
        <div>
          <img src={image1}/>
        </div>
      </CSSTransition>;
    return (
      <div>
        <div styleName="Desktop">
          <CSSTransition classNames="AppCollection" in appear={true} timeout={1000}>
            <div styleName="AppCollection">
              {icon}
            </div>
          </CSSTransition>
        </div>
        <input type="button" value="App"/>
        <input type="button" value="Desktop"/>
      </div>
    );
  }
}

export default CSSModules(AnimationTest, styles);
