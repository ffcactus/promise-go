import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import CircularProgress from '@material-ui/core/CircularProgress';
import styles from './App.css';


class EnclosureTask extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    if (this.props.task && this.props.task.Percentage !== 100) {
      return (
        <div styleName="list-progress">
          <CircularProgress size={28} variant="static" value={this.props.task.Percentage}/>
        </div>
      );
    }
    return <div/>;
  }
}

function mapStateToProps(state, props) {
  return { task: state.enclosureApp.enclosureTask.get(props.enclosureUri) };
}

EnclosureTask.propTypes = {
  enclosureUri: PropTypes.string,
  task: PropTypes.object,
};

export default connect(mapStateToProps)(CSSModules(EnclosureTask, styles));
