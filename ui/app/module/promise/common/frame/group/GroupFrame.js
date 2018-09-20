import React from 'react';
import SplitPane from 'react-split-pane';
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
          <SplitPane split="vertical" minSize={50} defaultSize="14.6%" resizerClassName={styles.Resizer + ' ' + styles.vertical}>
            {this.props.groupSection}
            <SplitPane split="vertical" minSize={50} defaultSize="27.6%" resizerClassName={styles.Resizer + ' ' + styles.vertical}>
              {this.props.listSection}
              {this.props.detailSection}
            </SplitPane>
          </SplitPane>
        </main>
        <footer styleName="footer border-column">
          {this.props.footer}
        </footer>
      </React.Fragment>
    );
  }
}

GroupFrame.propTypes = {
  styles: PropTypes.object,
  groupSection: PropTypes.object,
  listSection: PropTypes.object,
  detailSection: PropTypes.object,
  footer: PropTypes.object,
};

export default CSSModules(GroupFrame, styles, { allowMultiple: true });
