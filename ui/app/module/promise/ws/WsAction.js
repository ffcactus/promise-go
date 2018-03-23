function onmessage(messageString) {
  const message = JSON.parse(messageString);
  console.info(message);
}

/**
 * The action of create websocket connection.
 * @param {string} hostname The hostname to connect.
 */
function createWsConnection(hostname) {
  return (dispatch) => {
    const socket = new WebSocket('ws://' + hostname + '/promise/v1/ws');
    socket.onopen = () => {
      console.info('socket onopen.');
    };
    socket.onclose = () => {
      console.info('socket onclose');
    };
    socket.onmessage = (event) => {
      dispatch(onmessage(event.data));
    };
    socket.onerror = () => {
      console.info('socket onerror');
    };
  };
}

export { createWsConnection };
