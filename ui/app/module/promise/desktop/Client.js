import { doGet } from '../common/Client';

function getGlobal(hostname) {
  return doGet('http://' + hostname + '/promise/v1/global');
}

export { getGlobal };
