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
    this.groupResizerDragStart.bind(this);
    this.groupResizerDragover.bind(this);
    this.listResizerDragStart.bind(this);
  }

  groupResizerDragStart() {

  }

  groupResizerDragover(e) {
    const target = document.getElementById('group');
    const parentWidth = e.currentTarget.parentElement.clientWidth;
    const targetWitdh = e.clientX * 100 / parentWidth;
    const css = targetWitdh + '%';
    target.style.flexBasis = css;
  }

  listResizerDragover(e) {
    const target = document.getElementById('list');
    const parentWidth = e.currentTarget.parentElement.clientWidth;
    const targetWitdh = e.clientX * 100 / parentWidth;
    const css = targetWitdh + '%';
    target.style.flexBasis = css;
  }

  listResizerDragStart() {

  }

  render() {
    return (
      <React.Fragment>
        <header styleName="header">
          <section styleName="header-home">
            <p>Promise</p>
          </section>
        </header>
        <main styleName="main border-column" onDragOver={this.groupResizerDragover}>
          <section id="group" styleName="main-group">
            {this.props.groupSection}
          </section>
          <section id="group-resizer" styleName="main-resizer" draggble="true" onDragStart={this.listResizerDragStart}/>
          <section id="list" styleName="main-list">
            {this.props.listSection}
          </section>
          <section id="list-resizer" styleName="main-resizer" draggble="true" />
          <section styleName="main-detail">
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
