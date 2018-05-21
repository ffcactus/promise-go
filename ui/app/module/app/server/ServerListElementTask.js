import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import { Circle } from 'rc-progress';
import CenterDiv from '../../promise/common/CenterDiv';

class ServerListElementTask extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const task = this.props.serverTask.get(this.props.serverUri);
    if (task && task.Percentage !== 100) {
      return (
        <div styleName="ServerListElementTask">
          <CenterDiv>
            <Circle height="25px" trailColor="white" strokeWidth="12" percent={'' + task.Percentage} />
          </CenterDiv>
        </div>
      );
    }
    return <div styleName="ServerListElementTask"/>;
  }
}

function mapStateToProps(state) {
  return { serverTask: state.serverApp.serverTask };
}

ServerListElementTask.propTypes = {
  serverUri: PropTypes.string,
  serverTask: PropTypes.object,
};


export default connect(mapStateToProps)(CSSModules(ServerListElementTask, styles));
