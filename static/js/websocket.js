import { displayUsers } from "./home.js";

export let ws;

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
    console.log("Message from server ", event.data);

    try {
      const data = JSON.parse(event.data);
      console.log(1, data);
      // Si le message contient une information de connexion/déconnexion
      if (data.type === "log") {
        console.log("User connection update:", data.connexion);
        // Mettre à jour immédiatement la liste des utilisateurs
        displayUsers();
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
