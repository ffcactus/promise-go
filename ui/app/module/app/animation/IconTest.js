import React from 'react';
import CSSModules from 'react-css-modules';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { CSSTransition } from 'react-transition-group';
import styles from './Animation.css';
import IconNotification from '../../promise/common/IconNotification';

class IconTest extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const iconStyles = {
      appear: this.props.styles.IconAppear,
      appearActive: this.props.styles.IconAppearActive,
      enter: this.props.styles.IconEnter,
      enterActive: this.props.styles.IconEnterActive,
      enterDone: this.props.styles.IconEnterDone,
    };
    return (
      <div styleName="IconContainer">
        <CSSTransition classNames={iconStyles} in appear timeout={500}>
          <div key={this.props.name}>
            <div styleName="AppIconAndNotification">
              <Link to={'/xxxx'}>
                <img src={this.props.img} />
                <IconNotification notificationCount={this.props.notificationCount} />
              </Link>
            </div>
            <div styleName="AppIconName">
              <p>{this.props.name}</p>
            </div>
          </div>
        </CSSTransition>
      </div>
    );
  }
}

IconTest.propTypes = {
  img: PropTypes.string,
  name: PropTypes.string,
  notificationCount: PropTypes.number,
  styles: PropTypes.object,
};

export default CSSModules(IconTest, styles);
