import { sendPrivateMessage, ws } from "./websocket.js";

const unreadMessages = new Map();

export async function chatBox(nickname) {
  // Effacer les notifications
  clearNotifications(nickname);

  const app = document.getElementById("users");
  app.innerHTML = `
      <h2>Chat avec ${nickname}</h2>
      <div id="chatBox"></div>
      <input type="text" id="messageInput" placeholder="Votre message...">
      <button id="sendMessage">Envoyer</button>
    `;
  const sendMessageButton = document.getElementById("sendMessage");
  const messageInput = document.getElementById("messageInput");
  const chatBox = document.getElementById("chatBox");
  chatBox.dataset.nickname = nickname;

  sendMessageButton.addEventListener("click", () => {
    const message = messageInput.value;
    if (message) {
      sendPrivateMessage(nickname, message);
      chatBox.innerHTML += `<p><strong>Vous:</strong> ${message}</p><br>
      <span class="date">${new Date().toLocaleString()}</span>`;
      chatBox.scrollTop = chatBox.scrollHeight; // Scroll to the bottom
      messageInput.value = ""; // Clear the input after sending
    }
  });

  fetchMessages(nickname)
    .then((messages) => {
      messages.forEach((message) => {
        chatBox.innerHTML += `<p><strong>${message.sender}:</strong> ${
          message.content
        }</p><br>
        <span class="date">${new Date(
          message.created_at
        ).toLocaleString()}</span>`;
      });
    })
    .catch((error) => {
      console.error("Erreur lors de la récupération des messages", error);
    });
}

async function fetchMessages(nickname) {
  const response = await fetch(`/messages/${nickname}`);
  if (response.ok) {
    const messages = await response.json();
    return messages;
  } else {
    console.error("Erreur lors de la récupération des messages");
    return [];
  }
}

export function incrementUnreadMessages(sender) {
  const count = unreadMessages.get(sender) || 0;
  unreadMessages.set(sender, count + 1);
}

export function updateNotificationBadge(sender) {
  const userDiv = document.querySelector(
    `.users_user .nickname[data-nickname="${sender}"]`
  )?.parentElement;
  if (!userDiv) return;

  let badge = userDiv.querySelector(".notification-badge");
  const count = unreadMessages.get(sender) || 0;

  if (!badge) {
    badge = document.createElement("div");
    badge.className = "notification-badge";
    userDiv.appendChild(badge);
  }

  badge.textContent = count;
  badge.classList.add("active");
  badge.classList.add("pulse");
  setTimeout(() => badge.classList.remove("pulse"), 500);
}

function clearNotifications(sender) {
  unreadMessages.delete(sender);
  const userDiv = document.querySelector(
    `.users_user .nickname[data-nickname="${sender}"]`
  )?.parentElement;
  if (!userDiv) return;

  const badge = userDiv.querySelector(".notification-badge");
  if (badge) {
    badge.classList.remove("active");
  }
}
