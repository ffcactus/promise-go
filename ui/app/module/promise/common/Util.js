export function getParameterByName(name, url) {
  const _name = name.replace(/[\[\]]/g, '\\$&');
  const regex = new RegExp('[?&]' + _name + '(=([^&#]*)|&|#|$)');
  const results = regex.exec(url);
  if (!results) {
    return null;
  }
  if (!results[2]) {
    return '';
  }
  return decodeURIComponent(results[2].replace(/\+/g, ' '));
}
