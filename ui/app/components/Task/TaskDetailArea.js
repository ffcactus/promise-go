import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from '../../styles/Task/Task.css';

function TaskDetailArea(props) {
  return (
    <div styleName="TaskDetailArea">
      {props.task.current}
    </div>
  );
}

TaskDetailArea.propTypes = {
  task: PropTypes.object,
  dispatch: PropTypes.func,
};

function mapStateToProps(state) {
  const { task } = state;
  return { task };
}


export default connect(mapStateToProps)(CSSModules(TaskDetailArea, styles));
