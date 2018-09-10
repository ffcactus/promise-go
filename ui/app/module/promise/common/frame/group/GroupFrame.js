import React from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './GroupFrame.css';

/**
 * GroupFrame represents the 3 elements in body, the <header> <main> and the <footer>.
 */
class GroupFrame extends React.Component {
  constructor(props) {
    super(props);
    this.mainDragover.bind(this);
    this.groupResizerDragStart.bind(this);
    this.listResizerDragStart.bind(this);
  }

  mainDragover(e) {
    let target;
    e.preventDefault();
    console.info(e.target.getAttribute('data-id'));
    switch (e.target.dataset.text) {
      case 'group-resize':
        target = document.getElementById('group');
        break;
      case 'list-resize':
        target = document.getElementById('list');
        break;
      default:
        return;
    }
    const parentWidth = e.currentTarget.parentElement.clientWidth;
    const targetWitdh = e.clientX * 100 / parentWidth;
    const css = targetWitdh + '%';
    target.style.flexBasis = css;
  }

  listResizerDragStart() {
  }

  groupResizerDragStart() {
  }

  render() {
    return (
      <React.Fragment>
        <header styleName="header">
          <section styleName="header-home">
            <p>Promise</p>
          </section>
        </header>
        <main data-id="main" styleName="main border-column" onDragOver={this.mainDragover}>
          <section id="group" styleName="main-group">
            {this.props.groupSection}
          </section>
          <section data-id="group-resizer" id="group-resizer" styleName="main-resizer" draggble="true" onDragStart={()=>{}}/>
          <section id="list" styleName="main-list">
            {this.props.listSection}
          </section>
          <section data-id="list-resizer" id="list-resizer" styleName="main-resizer" draggble="true" onDragStart={this.listResizerDragStart}/>
          <section id="detail" styleName="main-detail">
            {this.props.detailSection}
          </section>
        </main>
        <footer styleName="footer border-column">
          {this.props.footer}
        </footer>
      </React.Fragment>
    );
  }
}

GroupFrame.propTypes = {
  groupSection: PropTypes.object,
  listSection: PropTypes.object,
  detailSection: PropTypes.object,
  footer: PropTypes.object,
};

export default CSSModules(GroupFrame, styles, { allowMultiple: true });
