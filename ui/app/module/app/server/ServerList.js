import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import ServerListElement from './ServerListElement';
import styles from './Server.css';

class ServerList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="ServerList">
        {
          this.props.server.serverList.length === 0 && <p>Empty</p>
        }
        {
          (this.props.server.serverList.length !== 0) && this.props.server.serverList.map((each) => {
            return (<ServerListElement key={each.URI} serverUri={each.URI}>{each.Name}</ServerListElement>);
          })
        }
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { server} = state;
  return { server };
}

ServerList.propTypes = {
  server: PropTypes.object,
  dispatch: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(ServerList, styles));

