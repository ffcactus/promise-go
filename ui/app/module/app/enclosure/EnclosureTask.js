import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import CircularProgress from '@material-ui/core/CircularProgress';
import styles from './App.css';


class EnclosureTask extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="list-progress">
        <CircularProgress size={28} variant="static" value={this.props.percentage}/>
      </div>
    );
  }
}


EnclosureTask.propTypes = {
  percentage: PropTypes.number.isRequired
};

export default CSSModules(EnclosureTask, styles);
