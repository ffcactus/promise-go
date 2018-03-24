import { doGet } from '../../../client/common';

export function getServerGroupList(hostname) {
  return doGet('http://' + hostname + '/promise/v1/servergroup');
}

export function getServerList(hostname) {
  return doGet('http://' + hostname + '/promise/v1/server');
}

export function getServer(hostname, uri) {
  return doGet('http://' + hostname + uri);
}
