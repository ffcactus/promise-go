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

function doPost(url, request) {
  return new Promise((resolve, reject) => {
    const xhr = createCORSRequest('POST', url, true);
    if (!xhr) {
      throw new Error('CORS not supported');
    }
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.onload = () => {
      resolve({
        response: JSON.parse(xhr.response),
        status: xhr.status,
        statusText: xhr.statusText
      });
    };
    xhr.onerror = (e) => {
      reject(e);
    };
    xhr.send(JSON.stringify(request));
  });
}

function doGet(url, request) {
  return new Promise((resolve, reject) => {
    // TODO do the check at a global place?
    const xhr = createCORSRequest('GET', url, true);
    if (!xhr) {
      throw new Error('CORS not supported');
    }
    xhr.setRequestHeader('Accept', 'application/json');
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.onload = () => {
      resolve({
        response: JSON.parse(xhr.response),
        status: xhr.status,
        statusText: xhr.statusText
      });
    };
    xhr.onerror = (e) => {
      reject(e);
    };
    xhr.send(JSON.stringify(request));
  });
}

export function login(hostname, userName, password) {
  return doPost('http://' + hostname + ':8080/director/rich/v1/login', {// + store.getState().global.host, {
    userName,
    password,
    domain: 'LOCAL'
  });
}

export function getServerList(hostname, start, count) {
  let uri = 'http://' + hostname + ':8080/director/rich/v1/server';
  if (start && count) {
    uri += ('?start=' + start + '&count=' + count);
  } else if (start) {
    uri += ('?start=' + start);
  } else if (count) {
    uri += ('?count=' + count);
  }
  return doGet(uri);
}


export function getTaskList(hostname, start, count) {
  let uri = 'http://' + hostname + ':8081/director/rich/v1/task';
  if (start && count) {
    uri += ('?start=' + start + '&count=' + count);
  } else if (start) {
    uri += ('?start=' + start);
  } else if (count) {
    uri += ('?count=' + count);
  }
  return doGet(uri);
}

export function loadServer(hostname, serverUri) {
  return doGet('http://' + hostname + ':8080' + serverUri);
}

export function getTask(hostname, taskUri) {
  return doGet('http://' + hostname + ':8081' + taskUri);
}
