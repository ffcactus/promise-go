import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import Styles from './AppFrame.css';

function GroupCollection(props) {
  return (
    <div styleName="GroupCollection">
      <div styleName="Group">
        {props.group}
      </div>
      <div styleName="Elements">
        {props.elements}
      </div>
      <div styleName="Detail">
        {props.detail}
      </div>
    </div>
  );
}

GroupCollection.propTypes = {
  group: PropTypes.object,
  elements: PropTypes.object,
  detail: PropTypes.object
};

export default CSSModules(GroupCollection, Styles);
