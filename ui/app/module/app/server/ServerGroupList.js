import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerGroupElement from './ServerGroupElement';

class ServerGroupList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="ServerGroupList">
        {this.props.server.serverGroupList.map((each) => {
          return <ServerGroupElement key={each.Name} name={each.Name} />;
        })}
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { server } = state;
  return { server };
}

ServerGroupList.propTypes = {
  server: PropTypes.object,
};

export default connect(mapStateToProps)(CSSModules(ServerGroupList, styles));

