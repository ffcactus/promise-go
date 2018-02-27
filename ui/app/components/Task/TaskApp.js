import React from 'react';
import CSSModules from 'react-css-modules';
import styles from '../../styles/Task/Task.css';
import TaskSearchArea from './TaskSearchArea';
import TaskListArea from './TaskListArea';
import TaskDetailArea from './TaskDetailArea';

const TaskApp = () =>
  <div styleName="TaskApp">
    <TaskSearchArea/>
    <TaskListArea/>
    <TaskDetailArea/>
  </div>;

export default CSSModules(TaskApp, styles);
