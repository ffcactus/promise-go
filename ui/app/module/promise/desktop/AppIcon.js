import React from 'react';
import CSSModules from 'react-css-modules';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { CSSTransition } from 'react-transition-group';
import IconNotification from './IconNotification';
import styles from './Desktop.css';


class AppIcon extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const iconStyles = {
      appear: this.props.styles.AppIconAppear,
      appearActive: this.props.styles.AppIconAppearActive,
      enter: this.props.styles.AppIconEnter,
      enterActive: this.props.styles.AppIconEnterActive,
      enterDone: this.props.styles.AppIconEnterDone,
    };
    return (
      <CSSTransition classNames={iconStyles} in appear timeout={300}>
        <div styleName="AppIconContainer">
          <div styleName="AppIconColumn">
            <div styleName="AppIconAndNotification">
              <Link to={this.props.uri}>
                <img src={this.props.img} />
                <IconNotification notificationCount={this.props.notificationCount} />
              </Link>
            </div>
            <div styleName="AppIconName">
              <p>{this.props.name}</p>
            </div>
          </div>
        </div>
      </CSSTransition>
    );
  }
}

AppIcon.propTypes = {
  uri: PropTypes.string,
  img: PropTypes.string,
  name: PropTypes.string,
  notificationCount: PropTypes.number,
  styles: PropTypes.object,
};

export default CSSModules(AppIcon, styles);
