import React from 'react';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import { CSSTransition } from 'react-transition-group';
import styles from './Phone.css';
import AppIcon from '../../promise/common/AppIcon';

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
      <CSSTransition classNames={iconStyles} in appear timeout={500}>
        <AppIcon
          key={this.props.name}
          name={this.props.name}
          image={this.props.img}
          uri={this.props.name}
          notificationCount={0} />
      </CSSTransition>
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
