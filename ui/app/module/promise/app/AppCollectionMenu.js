import React from 'react';
import PropTypes from 'prop-types';
import { push } from 'react-router-redux';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';

class AppCollectionMenu extends React.Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick(event) {
    event.preventDefault();
    this.props.dispatch(push('/'));
  }

  render() {
    return (
      <div styleName="AppCollectionMenu" onClick={this.handleClick}>
        <p styleName="AppCollectionMenuText">Promise</p>
      </div>
    );
  }
}

AppCollectionMenu.propTypes = {
  dispatch: PropTypes.func
};


export default connect()(CSSModules(AppCollectionMenu, Styles));
