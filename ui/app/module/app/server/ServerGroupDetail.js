import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import ServerGroupElement from './ServerGroupElement';

class ServerGroupDetail extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <React.Fragment>
        {this.props.server.groupList.map((each) => {
          return <ServerGroupElement key={each.Name} name={each.Name} />;
        })}
      </React.Fragment>
    );
  }
}

function mapStateToProps(state) {
  const { server } = state;
  return { server };
}

ServerGroupDetail.propTypes = {
  server: PropTypes.object,
};

export default connect(mapStateToProps)(ServerGroupDetail);

