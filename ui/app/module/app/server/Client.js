import { doGet, doPost } from '../../promise/common/Client';

export function getServerListByGroup(hostname, id) {
  return doGet('http://' + hostname + '/promise/v1/server-servergroup?$filter=ServerGroupID eq \'' + id + '\'');
}

export function getServerList(hostname) {
  return doGet('http://' + hostname + '/promise/v1/server');
}

export function getServer(hostname, uri) {
  return doGet('http://' + hostname + uri);
}

export function postServerGroup(hostname, dto) {
  return doPost('http://' + hostname + '/promise/v1/servergroup', dto);
}
