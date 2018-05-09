import React from 'react';
import PropTypes from 'prop-types';
import ServerDetailSectionPower from './ServerDetailSectionPower';

function ServerDetailTabChassis(props) {
  let content = null;
  if (props.chassis === null) {
    content = <p>Empty</p>;
  } else {
    content = (<div>
      <ServerDetailSectionPower power={props.chassis.Power} />
      <hr />
    </div>);
  }
  return content;
}

ServerDetailTabChassis.propTypes = {
  chassis: PropTypes.object
};

export default ServerDetailTabChassis;
