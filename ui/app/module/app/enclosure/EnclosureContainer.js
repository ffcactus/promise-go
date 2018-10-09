import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import Enclosure from './Enclosure';
import { AppState } from './ConstValue';
import * as EnclosureAppAction from './EnclosureAppAction';

class EnclosureContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    this.props.dispatch(EnclosureAppAction.appInit());
  }

  render() {
    switch (this.props.appState) {
      case AppState.LOADING:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
      case AppState.NORMAL:
        return <Enclosure />;
      case AppState.FAILURE:
        return <CenterDiv><p>App initialization failure.</p></CenterDiv>;
      default:
        return <CenterDiv><LoadingIcon /></CenterDiv>;
    }
  }
}

function mapStateToProps(state) {
  return {
    appState: state.enclosureApp.appState
  };
}

EnclosureContainer.propTypes = {
  appState: PropTypes.string,
  location: PropTypes.object,
  dispatch: PropTypes.func
};

export default connect(mapStateToProps)(EnclosureContainer);
