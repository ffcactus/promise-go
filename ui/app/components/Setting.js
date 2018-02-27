import React, {  Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Frame from './common/Frame';
import * as Action from '../actions/SettingAction';

class Setting extends Component {
  constructor(props) {
    super(props);
    this.createCoreRequest = this.createCoreRequest.bind(this);
    this.onUploadUpgradeBundleDialogOpen = this.onUploadUpgradeBundleDialogOpen.bind(this);
    this.onUploadUpgradeBundleConfirm = this.onUploadUpgradeBundleConfirm.bind(this);
  }

  createCoreRequest(method, url, async) {
    let xhr = new XMLHttpRequest();
    if ('withCredentials' in xhr) {
      // Check if the XMLHttpRequest object has a 'withCredentials' property.
      // 'withCredentials' only exists on XMLHTTPRequest2 objects.
      xhr.open(method, url, async);
    } else if (typeof XDomainRequest !== 'undefined') {
      // Otherwise, check if XDomainRequest.
      // XDomainRequest only exists in IE, and is IE's way of making CORS requests.
      xhr = new XDomainRequest();
      xhr.open(method, url, async);
    } else {
      // Otherwise, CORS is not supported by the browser.
      xhr = null;
    }
    return xhr;
  }

  onUploadUpgradeBundleDialogOpen(event) {
    event.preventDefault();
    if (window.File && window.FileReader && window.FileList && window.Blob) {
      this.props.dispatch(Action.settingUploadUpgradeBundleAction(document.getElementById('setting_upload_upgrade_bundle_input')));
    }
    // TODO else.
  }

  onUploadUpgradeBundleConfirm(event) {
    event.preventDefault();
    const fd = new FormData();
    const xhr = this.createCoreRequest('POST', 'http://192.168.116.135/rest/login', true); // asynchronously.
    if (!xhr) {
      throw new Error('CORS not supported');
    }

    xhr.onreadystatechange = () => {
      if (xhr.readyState === 4 && xhr.status === 200) {
        console.log(xhr.responseText);
      }
    };

    xhr.onloadstart = () => {
      console.log('onloadstart()');
    };

    xhr.onloadend = () => {
      console.log('onloadend()');
    };

    xhr.onload = () => {
      const responseText = xhr.responseText;
      console.log('onload() ' + responseText);
    };

    xhr.onerror = (e) => {
      console.log('onerror() There was an error! ' + e);
    };

    xhr.onprogress = (e) => {
      console.log('onprogress()');
      if (e.lengthComputable) {
        const percentage = Math.round((e.loaded * 100) / e.total);
        console.log('onprogress() ' + percentage);
      }
    };

    fd.append('file', this.props.setting.upgradeBundle);
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.send(JSON.stringify({
      userName: 'Administrator',
      domain: 'LOCAL',
      password: 'admin'
    }));
  }

  render() {
    const file = this.props.setting.upgradeBundle;

    const getMainContent = () => {
      return (
        <div>
          <label>
            Upload upgrade bundle<br />
            <input id="setting_upload_upgrade_bundle_input" type="file" onChange={this.onUploadUpgradeBundleDialogOpen}></input>
          </label>
          <ul>
            <li key="xxx">
              <strong>{file.name + ' '}</strong>
              {
                (file.type || 'n/a') + ' - ' +
                file.size + ' bytes, last modified: ' +
                (file.lastModifiedDate ? file.lastModifiedDate.toLocaleDateString() : 'n/a')
              }
            </li>
          </ul>
          <button onClick={this.onUploadUpgradeBundleConfirm}>Upload</button>
        </div>
      );
    };

    return (
      <Frame main={getMainContent()} footer={<p>footer</p>} />
    );
  }
}

function mapStateToProps(state) {
  const { setting } = state;
  return { setting };
}

Setting.propTypes = {
  setting: PropTypes.object,
  dispatch: PropTypes.func,
  history: PropTypes.object,
};

export default connect(mapStateToProps)(Setting);
