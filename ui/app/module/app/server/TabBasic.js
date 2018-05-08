import React from 'react';
import PropTypes from 'prop-types';
import SectionBasic from './SectionBasic';

function TabBasic(props) {
  let content;
  if (props.computerSystem === null) {
    content = <p>Empty</p>;
  } else {
    content = (<div>
      <SectionBasic server={props.server}/>
    </div>);
  }
  return content;
}

TabBasic.propTypes = {
  server: PropTypes.object
};

export default TabBasic;
