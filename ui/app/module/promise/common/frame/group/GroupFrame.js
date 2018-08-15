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
  }

  render() {
    return (
      <React.Fragment>
        <header styleName="header">
          <section styleName="header-home">
            <p>Promise</p>
          </section>
        </header>
        <main styleName="main border-column">
          <section id="group" styleName="main-group">
            {this.props.groupSection}
          </section>
          <section id="group-resizer" styleName="main-resizer" draggble />
          <section styleName="main-list">
            {this.props.listSection}
          </section>
          <section id="list-resizer" styleName="main-resizer" draggble />
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
