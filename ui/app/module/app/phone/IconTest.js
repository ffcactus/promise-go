import React from 'react';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import { CSSTransition } from 'react-transition-group';
import styles from './Phone.css';

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
        <div>
          <img src={this.props.img}/>
        </div>
      </CSSTransition>
    );
  }
}

IconTest.propTypes = {
  img: PropTypes.string,
  styles: PropTypes.object,
};

export default CSSModules(IconTest, styles);
