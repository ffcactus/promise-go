import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/Task/Task.css';
import TaskSimpleProgress from './TaskSimpleProgress';

function TaskListItem(props) {
  const task = props.task;
  return (
    <div styleName="TaskListItem">
      <div styleName="Name">{task.Name}</div>
      <div styleName="Description">{task.Description}</div>
      <div styleName="CurrentStep">{task.CurrentStep}</div>
      <div styleName="ProgressBar">
        <TaskSimpleProgress percentage={task.Percentage} />
      </div>
      <div styleName="Expend">Expend</div>
    </div>
  );
}

TaskListItem.propTypes = {
  task: PropTypes.object,
};

export default CSSModules(TaskListItem, styles);
