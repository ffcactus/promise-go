import React from 'react';
import PropTypes from 'prop-types';

class Progress extends React.Component {
  static propTypes = {
    completed: PropTypes.number,
    color: PropTypes.string,
    animation: PropTypes.number,
    height: PropTypes.oneOfType([
      PropTypes.string,
      PropTypes.number
    ])
  }

  static defaultProps = {
    completed: 0,
    color: '#0BD318',
    animation: 200,
    height: 10
  }

  static throwError() {
    return new Error(...arguments);
  }

  render() {
    const {color, completed, animation, height} = this.props;
    const style = {
      backgroundColor: color,
      width: completed + '%',
      transition: `width ${animation}ms`,
      height: height
    };

    return (
      <div className="progressbar-container">
        <div className="progressbar-progress" style={style}/>
      </div>
    );
  }
}

export default Progress;
