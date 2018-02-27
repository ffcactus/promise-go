import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from '../../styles/Task/Task.css';
import TaskSimple from './TaskSimple';

function TaskListArea(props) {
  return (
    <div styleName="TaskListArea">
      {props.task.taskList.map((each) => {
        return (
          <TaskSimple task={each} key={each.Uri} />
        );
      })}
    </div>
  );
}

TaskListArea.propTypes = {
  task: PropTypes.object,
  dispatch: PropTypes.func,
};

function mapStateToProps(state) {
  const { task } = state;
  return { task };
}


export default connect(mapStateToProps)(CSSModules(TaskListArea, styles));
