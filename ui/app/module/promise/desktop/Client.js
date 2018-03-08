import { doGet } from '../../../client/common';

function getGlobal(hostname) {
  return doGet('http://' + hostname + '/promise/v1/global');
}

export { getGlobal };
