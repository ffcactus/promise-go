import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from '../../styles/ServerFrame.css';
import TabSystem from './TabSystem';
import TabChassis from './TabChassis';
import TabBasic from './TabBasic';
import Tab from '../common/Tab';

function DetailFrame(props) {
  let content;
  if (props.server === null) {
    content = <p>Empty</p>;
  } else {
    const pages = [
      {
        'title': 'Basic',
        'content': <TabBasic server={props.server} />
      },
      {
        'title': 'System',
        'content': <TabSystem computerSystem={props.server.ComputerSystem}/>
      },
      {
        'title': 'Chassis',
        'content': <TabChassis chassis={props.server.Chassis}/>
      }
    ];
    content = <Tab pages={pages} />;
  }
  return (
    <div styleName="serverDetailFrame">
      {content}
    </div>
  );
}

DetailFrame.propTypes = {
  server: PropTypes.object
};

export default CSSModules(DetailFrame, styles);

