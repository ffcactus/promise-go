import React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import CSSModules from 'react-css-modules';
import PropTypes from 'prop-types';
import styles from './ServerController.css';
import FullSizeDiv from '../../promise/common/Widget/FullSizeDiv';
import * as Action from './Action';

class ServerController extends React.Component {
  constructor(props) {
    super(props);
    this._onClick = this._onClick.bind(this);
    this.state = {
      serverIndex: 0
    };
    this.padZeroes = this.padZeroes.bind(this);
    this.handleIndexChange = this.handleIndexChange.bind(this);
  }

  padZeroes(number, length) {
    let myString = '' + number;
    while (myString.length < length) {
      myString = '0' + myString;
    }
    return myString;
  }

  handleIndexChange(event) {
    this.setState({ serverIndex: event.target.value });
  }

  _onClick(event) {
    event.preventDefault();
    this.setState((state) => {
      return {
        ...state,
        serverIndex: state.serverIndex + 1
      };
    });
    this.props.dispatch(Action.addServer('Mock ' + this.padZeroes(this.state.serverIndex, 5)));
  }

  render() {
    return (
      <div>
        <FullSizeDiv>
          <Link to="/">Home</Link>
          <br/>
          <h1>Server Controller</h1>
          <input type="text" value={this.state.serverIndex} onChange={this.handleIndexChange}/>
          <input type="button" value="Create" onClick={this._onClick}/>
          <h1>Response</h1>
          <section>
            <p>{JSON.stringify(this.props.response, null, 2)}</p>
          </section>
          <h1>Message</h1>
          <section>
            <p>{JSON.stringify(this.props.message, null, 2)}</p>
          </section>
          <h1>Exception</h1>
          <section>
            <p>{JSON.stringify(this.props.exception, null, 2)}</p>
          </section>
        </FullSizeDiv>
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { response, message, exception } = state.serverControllerApp;
  return {
    response,
    message,
    exception
  };
}

ServerController.propTypes = {
  dispatch: PropTypes.func,
  response: PropTypes.object,
  message: PropTypes.array,
  exception: PropTypes.any
};

export default connect(mapStateToProps)(CSSModules(ServerController, styles));
