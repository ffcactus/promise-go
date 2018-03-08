import { doPost } from '../../../client/common';

function login(hostname, userName, password) {
  return doPost('http://' + hostname + '/promise/v1/auth/login', {
    'Name': userName,
    'Password': password
  });
}

export { login };
