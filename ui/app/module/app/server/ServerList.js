import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import * as ServerAction from './ServerAction';
import ServerListElement from './ServerListElement';
import styles from './Server.css';

class ServerList extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    // At this point, the servergroup hasn't been clicked yet, so let's do a auto get collection.
    this.props.dispatch(ServerAction.onListDidMount());
  }

  render() {
    return (
      <div styleName="ServerList">
        {
          this.props.serverApp.serverList.size === 0 && <p>Empty</p>
        }
        {
          (this.props.serverApp.serverList.size !== 0) && this.props.serverApp.serverList.map((value, key) => {
            // we only have URI at this moment.
            return (<ServerListElement key={key} serverUri={key} />);
          }).toArray()
        }
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { serverApp } = state;
  return { serverApp, updated: true };
}

ServerList.propTypes = {
  serverApp: PropTypes.object,
  dispatch: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(ServerList, styles));

