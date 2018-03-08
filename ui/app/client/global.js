import { doGet } from './common';

export function getGlobal(hostname) {
  return doGet('http://' + hostname + '/promise/v1/global');
}
