// import store from '../configureStore';

const Promise = require('promise');
// const HOST = (process.env.NODE_ENV === 'development') ? '192.168.116.135' : store.getState().global.host;

function createCORSRequest(method, url, async) {
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

export function doPost(url, request) {
  return new Promise((resolve, reject) => {
    const xhr = createCORSRequest('POST', url, true);
    if (!xhr) {
      throw new Error('CORS not supported');
    }
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.onload = () => {
      if (xhr.status === 500) {
        reject('Internal Error.');
        return;
      }
      resolve({
        response: JSON.parse(xhr.response),
        status: xhr.status,
        statusText: xhr.statusText
      });
    };
    xhr.onabort = (e) => {
      reject(e);
    };
    xhr.ontimeout = (e) => {
      reject(e);
    };
    xhr.onerror = (e) => {
      reject(e);
    };
    xhr.send(JSON.stringify(request));
  });
}

export function doGet(url, request) {
  return new Promise((resolve, reject) => {
    // TODO do the check at a global place?
    const xhr = createCORSRequest('GET', url, true);
    if (!xhr) {
      throw new Error('CORS not supported');
    }
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.onload = () => {
      if (xhr.status === 500) {
        reject('Internal Error.');
        return;
      }
      resolve({
        response: JSON.parse(xhr.response),
        status: xhr.status,
        statusText: xhr.statusText
      });
    };
    xhr.onabort = (e) => {
      reject(e);
    };
    xhr.ontimeout = (e) => {
      reject(e);
    };
    xhr.onerror = (e) => {
      reject(e);
    };
    xhr.send(JSON.stringify(request));
  });
}

/**
 * GetAction is the common get action template.
 * @param {string} uri The URI to perform GET method.
 * @param {ActionType} start The ActionType when start.
 * @param {ActionType} success The ActionType when success.
 * @param {ActionType} message The ActionType when message returned.
 * @param {ActionType} exception The ActionType when exception returned.
 */
export function createGetAction(uri, start, success, message, exception) {
  return (dispatch, getState) => {
    dispatch({ type: start });
    doGet('http://' + getState().session.hostname + uri).then((resp) => {
      if (resp.status === 200) {
        dispatch({ type: success, info: resp.response });
        return;
      }
      if (resp.status === 400) {
        dispatch({ type: message, info: resp.response });
        return;
      }
    }).catch((e) => {
      dispatch({ type: exception, info: e });
    });
  };
}
