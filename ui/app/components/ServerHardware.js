import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Frame from './common/Frame';
import ServerListFrame from './Server/ServerListFrame';
import ServerListElement from './Server/ServerListElement';
import DetailFrame from './Server/DetailFrame';
import ServerAppFrame from './Server/ServerAppFrame';

// import * as Action from '../actions/HardwareAction';

class ServerHardware extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const getMainContent = () => {
      return (
        <ServerAppFrame>
          <ServerListFrame>
            {
              this.props.server.serverList.length === 0 && <p>Empty</p>
            }
            {
              (this.props.server.serverList.length !== 0) && this.props.server.serverList.map((each) => {
                return (<ServerListElement key={each.Uri} serverUri={each.Uri}>{each.Name}</ServerListElement>);
              })
            }
          </ServerListFrame>
          <DetailFrame server={this.props.server.current} />
        </ServerAppFrame>
      );
    };

    return (
      <Frame main={getMainContent()} footer={<p>footer</p>} />
    );
  }
}

function mapStateToProps(state) {
  const { server } = state;
  return { server };
}

ServerHardware.propTypes = {
  server: PropTypes.object,
  dispatch: PropTypes.func,
};

export default connect(mapStateToProps)(ServerHardware);
