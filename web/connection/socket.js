function setupWebSocket() {
  const socket = new WebSocket("ws://" + window.location.host + "/ws");

  socket.onmessage = function (event) {
    // Handle incoming messages from the WebSocket
    console.log(event)
  };
}

document.addEventListener("DOMContentLoaded", setupWebSocket);