import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import Styles from './GroupCollectionApp.css';

function GroupCollectionApp(props) {
  return (
    <div styleName="GroupCollectionApp">
      <div styleName="Group">
        {props.group}
      </div>
      <div styleName="ElementList">
        {props.elementList}
      </div>
      <div styleName="Detail">
        {props.detail}
      </div>
    </div>
  );
}

GroupCollectionApp.propTypes = {
  group: PropTypes.object,
  elementList: PropTypes.object,
  detail: PropTypes.object
};

export default CSSModules(GroupCollectionApp, Styles);
