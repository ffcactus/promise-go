import React from 'react';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

class Settings extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <React.Fragment>
        <ul>
          <li>
            <p>Username: {this.props.username}</p>
            <p>email: {this.props.email}</p>
          </li>
        </ul>
        <Link to="/">Home</Link>
      </React.Fragment>
    );
  }
}

function mapStateToProps(state) {
  return {
    username: state.session.username,
    email: state.session.email,
  };
}

Settings.propTypes = {
  username: PropTypes.string,
  email: PropTypes.string,
};

export default connect(mapStateToProps)(Settings);
