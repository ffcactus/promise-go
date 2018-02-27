import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/Task/Task.css';
import Progress from '../common/Progress';

function TaskSimpleProgress(props) {
  return (
    <Progress completed={props.percentage} height="5px"/>
  );
}

TaskSimpleProgress.propTypes = {
  percentage: PropTypes.number,
};

export default CSSModules(TaskSimpleProgress, styles);
