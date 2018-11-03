import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import { AutoSizer, List } from 'react-virtualized';
import { OrderedMap } from 'immutable';
import EnclosureListElement from './EnclosureListElement';
import styles from './App.css';

class EnclosureList extends React.Component {
  constructor(props) {
    super(props);
    this.rowRenderer = this.rowRenderer.bind(this);
    this.listRef = React.createRef();
  }

  componentWillReceiveProps() {
    this.listRef.current.forceUpdateGrid();
  }

  rowRenderer({key, index, style}) {
    const enclosure = this.props.enclosureOrderedMap.toIndexedSeq().get(index);
    return <EnclosureListElement key={key} enclosure={enclosure} style={style}/>;
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container left-border flex-item-last">
        <AutoSizer>{
          ({ height, width }) => (
            <List
              ref= {this.listRef}
              width={width}
              height={height}
              rowCount={this.props.enclosureOrderedMap.size}
              scrollToIndex={this.props.enclosureIndex}
              rowHeight={40}
              rowRenderer={this.rowRenderer}
            />
          )
        }</AutoSizer>
      </div>
    );
  }
}

EnclosureList.propTypes = {
  setListRef: PropTypes.func,
  enclosureIndex: PropTypes.number,
  enclosureOrderedMap: PropTypes.objectOf(OrderedMap),
};

function mapStateToProps(state) {
  return {
    enclosureOrderedMap: state.enclosureApp.enclosureOrderedMap
  };
}

export default connect(mapStateToProps)(CSSModules(EnclosureList, styles, {allowMultiple: true}));
