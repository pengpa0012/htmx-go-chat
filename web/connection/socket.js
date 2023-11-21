const socket = new WebSocket("ws://" + window.location.host + "/ws");

function setupWebSocket() {

  socket.onmessage = function (event) {
    // Handle incoming messages from the WebSocket
    console.log(event)
    newMessage(event)
  };

  const form = document.getElementById("form")
  form.addEventListener("submit", e => {
    e.preventDefault()
    console.log("test")
    socket.send("test")
  })
}


function newMessage (event) {
  console.log(event)
  const newMessage = event.data
  const chatDiv = document.querySelector(".chats");
  const messageElement = document.createElement("li");
  messageElement.textContent = newMessage;
  chatDiv.appendChild(messageElement);
}


document.addEventListener("DOMContentLoaded", setupWebSocket);