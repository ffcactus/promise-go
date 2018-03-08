import React from 'react';
import { Route } from 'react-router';
import App from '../components/App';
import About from '../components/About';
import Login from '../components/Login';
import StartupContainer from '../containers/StartupContainer';
import Desktop from '../components/Desktop/Desktop';
import DashBoard from '../components/DashBoard';
import TaskApp from '../components/Task/TaskApp';
import ServerHardware from '../components/ServerHardware';
import Setting from '../components/Setting';
import Example from '../components/Example';
import Platform from '../containers/PlatformContainer';

export default class RTRouter extends React.Component {
  constructor() {
    super();
  }

  render() {
    return (
      <div>
        <Route exact path="/" components={App} />
        <Route path="/login" component={Login} />
        <Route path="/desktop" component={Desktop} />
        <Route path="/startup" component={StartupContainer} />
        <Route path="/dashboard" component={DashBoard} />
        <Route path="/task" component={TaskApp} />
        <Route path="/server" component={ServerHardware} />
        <Route path="/setting" component={Setting} />
        <Route path="/example" component={Example} />
        <Route path="/about" component={About} />
        <p>Route</p>
      </div>
    );
  }
}
