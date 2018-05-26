import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerSortOrderSelect from './ServerSortOrderSelect';

class ServerListControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="ServerListControlArea" >
        <ServerSortOrderSelect listRef={this.props.listRef} options={['Name', 'Health']} />
      </div>
    );
  }
}

ServerListControlArea.propTypes = {
  listRef: PropTypes.object,
};

export default CSSModules(ServerListControlArea, styles);
