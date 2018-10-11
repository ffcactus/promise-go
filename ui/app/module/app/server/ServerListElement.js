import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import * as ServerAction from './ServerAction';
import ServerListElementTask from './ServerListElementTask';
import ServerListElementName from './ServerListElementName';
import { Health } from '../../promise/common/Widget/Health';
import CenterDiv from '../../promise/common/CenterDiv';

class ServerListElement extends React.PureComponent {
  constructor(props) {
    super(props);
    this.onSelect = this.onSelect.bind(this);
  }

  // On selecting we need display the detail infomation.
  onSelect(event) {
    event.preventDefault();
    this.props.dispatch(ServerAction.uiListSelect(this.props.server.URI));
  }

  render() {
    const currentStyle = 'ServerListElement ' + (
      this.props.selected ?
        'Selected' : 'NotSelected'
    );
    return (
      <div styleName={currentStyle} onClick={this.onSelect}>
        <div styleName="ServerListElementHealth">
          <CenterDiv>
            <Health health={this.props.server.Health}/>
          </CenterDiv>
        </div>
        <ServerListElementName name={this.props.server.Name ? this.props.server.Name : '...'} />
        <ServerListElementTask serverUri={this.props.server.URI}/>
      </div>
    );
  }
}

function mapStateToProps(state, ownProps) {
  return {
    selected: ownProps.server.URI === state.serverApp.currentServerUri,
  };
}

ServerListElement.propTypes = {
  server: PropTypes.object,
  selected: PropTypes.bool,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(CSSModules(ServerListElement, styles, {allowMultiple: true}));

