import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerGroupElement from './ServerGroupElement';
import * as ServerGroupAction from './ServerGroupAction';

class ServerGroupList extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    this.props.dispatch(ServerGroupAction.getCollection());
  }

  render() {
    return (
      <div styleName="ServerGroupList">
        {this.props.serverApp.serverGroupList.map((each) => {
          return <ServerGroupElement key={each.ID} servergroup={each} />;
        })}
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { serverApp } = state;
  return { serverApp };
}

ServerGroupList.propTypes = {
  dispatch: PropTypes.func,
  serverApp: PropTypes.object,
};

export default connect(mapStateToProps)(CSSModules(ServerGroupList, styles));

