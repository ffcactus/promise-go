import React, {  Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import Frame from './common/Frame';
import Styles from '../styles/DashBoard.css';

class DashBoard extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    let red = 0;
    let yellow = 0;
    let green = 0;
    let gray = 0;
    for (let i = 0; i < this.props.server.serverList.length; i++) {
      switch (this.props.server.serverList[i].Health) {
        case 'OK':
          green++;
          break;
        case 'Warning':
          yellow++;
          break;
        case 'Critical':
          red++;
          break;
        default:
          gray++;
          break;
      }
    }
    const chartData = {
      datasets: [{
        data: [red, yellow, green, gray],
        backgroundColor: [
          'red',
          'yellow',
          'green',
          'gray'
        ]
      }],

      // These labels appear in the legend and in the tooltips when hovering different arcs
      labels: [
        '' + red,
        '' + yellow,
        '' + green,
        '' + gray
      ]
    };
    const chart = <div styleName="ChartArea"><Doughnut data={chartData} /><p>Server</p></div>;
    return <Frame main={chart} footer={<p>footer</p>} />;
  }
}

function mapStateToProps(state) {
  const { server } = state;
  return { server };
}

DashBoard.propTypes = {
  server: PropTypes.object,
  dispatch: PropTypes.func,
};

export default connect(mapStateToProps)(CSSModules(DashBoard, Styles));
