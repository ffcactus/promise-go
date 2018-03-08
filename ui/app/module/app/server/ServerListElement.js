import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
// import * as Action from '.ServerAction';

class ServerListElement extends React.Component {
  constructor(props) {
    super(props);
    this.onSelect = this.onSelect.bind(this);
    this.state = {
      selected: false
    };
  }

  shouldComponentUpdate(nextProps, nextState) {
    return this.state.selected !== nextState.selected;
  }

  onSelect(event) {
    event.preventDefault();
    // this.props.dispatch(Action.loadServer(this.props.serverUri));
  }

  render() {
    return (
      <div styleName="ServerListElement" onClick={this.onSelect}>{this.props.children}</div>
    );
  }
}

function mapStateToProps(state) {
  const { server } = state;
  return { server };
}

ServerListElement.propTypes = {
  serverUri: PropTypes.string,
  children: PropTypes.string,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(CSSModules(ServerListElement, styles));

