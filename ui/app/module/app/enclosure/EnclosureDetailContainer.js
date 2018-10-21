import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as EnclosureAction from './EnclosureAction';
import { ResourceDetailState } from './ConstValue';
import CenterDiv from '../../promise/common/CenterDiv';
import LoadingIcon from '../../promise/common/LoadingIcon';
import EnclosureDetail from './EnclosureDetail';

class EnclosureDetailContainer extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
    if (this.props.enclosure === null && this.props.enclosureUri !== null) {
      this.props.dispatch(EnclosureAction.selectElement(this.props.enclosureUri));
    }
  }

  render() {
    switch(this.props.resourceDetailState) {
      case ResourceDetailState.EMPTY:
        return <CenterDiv><p>No Enclosure Selected.</p></CenterDiv>;
      case ResourceDetailState.LOADING:
        return <CenterDiv><LoadingIcon/></CenterDiv>;
      case ResourceDetailState.READY:
        return <EnclosureDetail enclosure={this.props.enclosure} />;
      case ResourceDetailState.FAILURE:
        return <CenterDiv><p>Loading Enclosure Failed</p></CenterDiv>;
      default:
        return <CenterDiv><LoadingIcon/></CenterDiv>;
    }
  }
}

EnclosureDetailContainer.propTypes = {
  enclosure: PropTypes.object,
  enclosureUri: PropTypes.string,
  resourceDetailState: PropTypes.string,
  dispatch: PropTypes.func,
};

function mapStateToProps(state) {
  return {
    enclosureUri: state.serverApp.enclosureUri,
    enclosure: state.enclosureApp.enclosure,
    resourceDetailState: state.enclosureApp.resourceDetailState,
  };
}

export default connect(mapStateToProps)(EnclosureDetailContainer);
