import React from 'react';
import CSSModules from 'react-css-modules';
import { Link } from 'react-router';
import styles from './Menu.css';

const Menu = () =>
  <header styleName="header">
    <Link to="/" styleName="item">Dashboard</Link>
    <Link to="/desktop" styleName="item">Desktop</Link>
    <Link to="/startup" styleName="item">Startup</Link>
    <Link to="/task" styleName="item">Task</Link>
    <Link to="/server" styleName="item">Server</Link>
    <Link to="/login" styleName="item">Login</Link>
    <Link to="/setting" styleName="item">Setting</Link>
    <Link to="/example" styleName="item">Example</Link>
    <Link to="/about" styleName="item">About</Link>
  </header>;

export default CSSModules(Menu, styles);
