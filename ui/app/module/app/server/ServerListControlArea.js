import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerSortOrderSelect from './ServerSortOrderSelect';
import ServerControlAdd from './ServerControlAdd';
import ServerControlEdit from './ServerControlEdit';

class ServerListControlArea extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container border-column-first" style={{maxHeight: '40px'}}>
        <ServerSortOrderSelect listRef={this.props.listRef} options={['Name', 'Health']} />
        <ServerControlAdd />
        <ServerControlEdit />
      </div>
    );
  }
}

ServerListControlArea.propTypes = {
  listRef: PropTypes.object,
};

export default CSSModules(ServerListControlArea, styles, {allowMultiple: true});
