import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import * as ServerAction from './ServerAction';

class ServerListElement extends React.Component {
  constructor(props) {
    super(props);
    this.onSelect = this.onSelect.bind(this);
    this.state = {
      selected: false
    };
  }

  // shouldComponentUpdate(nextProps, nextState) {
  //   return this.state.selected !== nextState.selected;
  // }
  componentDidMount() {
    this.props.dispatch(ServerAction.getServer(this.props.serverUri));
  }

  onSelect(event) {
    event.preventDefault();
    // this.props.dispatch(Action.loadServer(this.props.serverUri));
  }

  render() {
    const server = this.props.serverApp.serverList.get(this.props.serverUri);

    return (
      <div styleName="ServerListElement" onClick={this.onSelect}>{server.Name}</div>
    );
  }
}

function mapStateToProps(state) {
  const { serverApp } = state;
  return { serverApp };
}

ServerListElement.propTypes = {
  serverApp: PropTypes.object,
  serverUri: PropTypes.string,
  children: PropTypes.string,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(CSSModules(ServerListElement, styles, {allowMultiple: true}));

