import { displayUsers } from "./home.js";
import { incrementUnreadMessages, updateNotificationBadge } from "./message.js";

export let ws = null;

const unreadMessages = new Map(); // Pour stocker le nombre de messages non lus par utilisateur

export function InitWS() {
  ws = new WebSocket("ws://localhost:8080/ws");

  ws.onopen = function (event) {
    console.log("WebSocket is open now.");
  };

  ws.onclose = function (event) {
    console.log("WebSocket is closed now.");
    console.log("Message from server ", event);
  };

  ws.onmessage = function (event) {
    try {
      const data = JSON.parse(event.data);
      if (data.type === "log") {
        console.log("User connection update:", data.connexion);
        displayUsers();
      } else if (data.type === "message") {
        const chatBox = document.getElementById("chatBox");
        if (chatBox && chatBox.dataset.nickname === data.sender) {
          chatBox.innerHTML += `<p><strong>${data.sender}:</strong> ${
            data.content
          }</p><br>
          <span class="date">${new Date(
            data.created_at
          ).toLocaleString()}</span>`;
        } else {
          // Incrémenter le compteur de messages non lus
          incrementUnreadMessages(data.sender);
          updateNotificationBadge(data.sender);
        }
      }
    } catch (e) {
      console.error("Error parsing message:", e);
    }
  };

  ws.onerror = (error) => {
    console.error("WebSocket error observed:", error);
    ws.close();
  };
}

export function sendPrivateMessage(nickname, message) {
  console.log(ws);
  if (ws && ws.readyState === WebSocket.OPEN) {
    ws.send(
      JSON.stringify({
        type: "message",
        receiver: nickname,
        content: message,
        created_at: new Date(),
      })
    );
  } else {
    console.error("WebSocket is not open. Cannot send message.");
  }
}
