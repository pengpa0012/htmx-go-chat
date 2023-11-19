function setupWebSocket() {
  const socket = new WebSocket("ws://" + window.location.host + "/ws");

  socket.onmessage = function (event) {
    // Handle incoming messages from the WebSocket
    console.log(event)
    const newMessage = JSON.parse(event.data);
    const chatDiv = document.getElementById("chats");
    const messageElement = document.createElement("p");
    messageElement.textContent = newMessage.Message;
    chatDiv.appendChild(messageElement);
  };
}

document.addEventListener("DOMContentLoaded", setupWebSocket);