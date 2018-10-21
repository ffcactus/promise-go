import React from 'react';
import PropTypes from 'prop-types';
import { push } from 'react-router-redux';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './AppFrame.css';

class AppCollectionMenu extends React.Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick(event) {
    event.preventDefault();
    this.props.dispatch(this.props.action);
    this.props.dispatch(push('/'));
  }

  render() {
    return (
      <div styleName="AppCollectionMenu" onClick={this.handleClick}>
        <p>Promise</p>
      </div>
    );
  }
}

AppCollectionMenu.propTypes = {
  dispatch: PropTypes.func,
  action: PropTypes.object,
};


export default connect()(CSSModules(AppCollectionMenu, styles));
