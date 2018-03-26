// function onServerMessage(message)
// {
//   switch(message.Type) {
//     case 'Create':
//       return {
//         type: ''
//       }
//   }
//   return {
//     type: 'xxxx'
//   };
// }

const handlerMap = new Map();

function onmessage(messageString) {
  const message = JSON.parse(messageString);
  const handler = handlerMap[message.Category];
  if (handler) {
    return handler(message);
  }
  return ()=>{};
}

function registerMessageAction(category, handler) {
  handlerMap[category] = handler;
}

function unregisterMessageAction(category) {
  return handlerMap.delete(category);
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

export { createWsConnection, registerMessageAction, unregisterMessageAction };
