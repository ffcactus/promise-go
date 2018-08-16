import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import styles from './Server.css';
import ServerGroupElement from './ServerGroupElement';

class ServerGroupList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    const size = this.props.serverGroupList.size;
    const list = [];
    for (let i = 0; i < size; i++) {
      const each = this.props.serverGroupList.get(i);
      list.push(<ServerGroupElement key={each.URI} servergroup={each}/>);
    }
    return <div styleName="flex-item-last border-column">{list}</div>;
  }
}

function mapStateToProps(state) {
  return { serverGroupList: state.serverApp.serverGroupList };
}

ServerGroupList.propTypes = {
  dispatch: PropTypes.func,
  serverGroupList: PropTypes.object,
};

export default connect(mapStateToProps)(CSSModules(ServerGroupList, styles, {allowMultiple: true}));

