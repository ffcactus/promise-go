import React from "react";
import PropTypes from "prop-types";
import CSSModules from "react-css-modules";
import styles from "./Tab.css";

class Tab extends React.Component {
  constructor(props) {
    super(props);
    this.onTabClick = this.onTabClick.bind(this);
  }

  componentDidMount() {
    const id = this.props.defaultOpen
      ? this.props.defaultOpen
      : this.props.pages[0].title;
    const defaultOpen = document.getElementById("button" + id);
    if (defaultOpen !== null) {
      defaultOpen.click();
    }
  }

  onTabClick(event) {
    let i;
    event.preventDefault();
    this.props.handler(event.currentTarget.value);
    const tabContent = document.getElementsByClassName("tabcontent");
    for (i = 0; i < tabContent.length; i++) {
      tabContent[i].style.display = "none";
    }
    const tabLinks = document.getElementsByClassName("tablinks");
    for (i = 0; i < tabLinks; i++) {
      tabLinks[i].className = tabLinks[i].className.replace(" active", "");
    }
    document.getElementById(event.currentTarget.value).style.display = "block";
    event.currentTarget.className += " active";
  }

  render() {
    return (
      <div styleName="TabRoot">
        <div styleName="TabHeader">
          {this.props.pages.map(each => {
            return (
              <button
                key={each.title}
                id={"button" + each.title}
                onClick={this.onTabClick}
                value={each.title}
              >
                {each.title}
              </button>
            );
          })}
        </div>
        {this.props.pages.map(each => {
          return (
            <div id={each.title} key={each.title} styleName="TabContent">
              {each.content}
            </div>
          );
        })}
      </div>
    );
  }
}

Tab.propTypes = {
  pages: PropTypes.arrayOf(PropTypes.object),
  handler: PropTypes.func,
  defaultOpen: PropTypes.string
};

export default CSSModules(Tab, styles);
