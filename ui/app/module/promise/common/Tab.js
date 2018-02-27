import React, {  Component } from 'react';
import PropTypes from 'prop-types';
import CSSModules from 'react-css-modules';
import styles from './Tab.css';

class Tab extends Component {
  constructor(props) {
    super(props);
    this.onTabClick = this.onTabClick.bind(this);
  }

  componentDidMount() {
    const defaultOpen = document.getElementById('button' + this.props.pages[0].title);
    if (defaultOpen !== null) {
      defaultOpen.click();
    }
  }

  onTabClick(event) {
    let i;
    event.preventDefault();
    const tabContent = document.getElementsByClassName('tabcontent');
    for (i = 0; i < tabContent.length; i++) {
      tabContent[i].style.display = 'none';
    }
    const tabLinks = document.getElementsByClassName('tablinks');
    for (i = 0; i < tabLinks; i++) {
      tabLinks[i].className = tabLinks[i].className.replace(' active', '');
    }
    document.getElementById(event.currentTarget.value).style.display = 'block';
    event.currentTarget.className += ' active';
  }

  render() {
    return (
      <div className={styles.root}>
        <div className="tab">
          {this.props.pages.map(each => {
            return (
              <button key={each.title} id={'button' + each.title} onClick={this.onTabClick} value={each.title}>{each.title}</button>
            );
          })}
        </div>
        {this.props.pages.map(each => {
          return (
            <div id={each.title} key={each.title} className="tabcontent">{each.content}</div>
          );
        })}
      </div>
    );
  }
}

Tab.propTypes = {
  pages: PropTypes.array,
};

export default CSSModules(Tab, styles);
