import { displayUsers } from "./home.js";

export function InitWS() {
  const ws = new WebSocket("ws://localhost:8080/ws");
  ws.onopen = function (event) {
    console.log("WebSocket is open now.");
  };
  ws.onclose = function (event) {
    console.log("WebSocket is closed now.");
    // if (event.code !== 1000) {
    //   console.error(
    //     `🔁 Tentative de reconnexion dans ${reconnectInterval / 1000}s...`
    //   );
    //   setTimeout(InitWS, reconnectInterval);
    // }
  };
  ws.onmessage = function (event) {
    console.log("Message from server ", event.data);

    const data = JSON.parse(event.data);
    if (data.type === "message") {
      const message = data.message;
      const chatBox = document.getElementById("chat-box");
      chatBox.innerHTML += `<p>${message}</p>`;
    } else if (data.connexion) {
      displayUsers();
    }
  };
  ws.onerror = (error) => {
    console.error("WebSocket error observed:", event);
    ws.close();
  };
}
