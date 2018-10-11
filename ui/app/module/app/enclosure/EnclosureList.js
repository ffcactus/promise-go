import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import CSSModules from 'react-css-modules';
import { AutoSizer, List } from 'react-virtualized';
import EnclosureListElement from './EnclosureListElement';
import styles from './Enclosure.css';

class EnclosureList extends React.Component {
  constructor(props) {
    super(props);
    this.rowRenderer = this.rowRenderer.bind(this);
  }

  rowRenderer({key, index}) {
    const enclosure = this.props.enclosureList.get(index);
    return <EnclosureListElement key={key} enclosure={enclosure} />;
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container border-column flex-item-last">
        <AutoSizer>{
          ({ height, width }) => (
            <List
              ref={this.props.setListRef}
              width={width}
              height={height}
              rowCount={this.props.enclosureList.size}
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
  enclosureList: PropTypes.object,
};

function mapStateToProps(state) {
  return {
    enclosureList: state.enclosureApp.enclosureList
  };
}

export default connect(mapStateToProps)(CSSModules(EnclosureList, styles, {allowMultiple: true}));
