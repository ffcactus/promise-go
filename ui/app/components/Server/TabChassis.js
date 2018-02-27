import React from 'react';
import PropTypes from 'prop-types';
import SectionPower from './SectionPower';

function TabChassis(props) {
  let content = null;
  if (props.chassis === null) {
    content = <p>Empty</p>;
  } else {
    content = (<div>
      <SectionPower power={props.chassis.Power} />
      <hr />
    </div>);
  }
  return content;
}

TabChassis.propTypes = {
  chassis: PropTypes.object
};

export default TabChassis;
