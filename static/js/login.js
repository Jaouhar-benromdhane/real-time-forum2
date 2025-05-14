import { InitWS } from "./websocket.js";

export function loadLogin() {
  const app = document.getElementById("app");
  app.innerHTML = `
      <h2>Connexion</h2>
      <form id="loginForm">
        <label for="identifiant">Identifiant (email ou pseudo) :</label><br>
        <input type="text" id="identifiant" name="identifiant" required><br><br>
  
        <label for="password">Mot de passe :</label><br>
        <input type="password" id="password" name="password" required><br><br>
  
        <button type="submit">Se connecter</button>
      </form>
      <div id="loginMessage"></div>
    `;

  const loginForm = document.getElementById("loginForm");

  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const identifiant = document.getElementById("identifiant").value;
    const password = document.getElementById("password").value;

    const response = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ identifiant, password }),
    });

    const messageDiv = document.getElementById("loginMessage");

    if (response.ok) {
      messageDiv.innerText = "Connexion réussie !";

      // Basculer l'affichage de la navigation
      document.querySelector(".online").style.display = "block";
      document.querySelector(".offline").style.display = "none";

      // Initialiser la connexion WebSocket
      InitWS();

      // Redirection vers la page d'accueil (SPA)
      navigateTo("home");
    } else {
      const errorText = await response.text();
      messageDiv.innerText = "Erreur : " + errorText;
    }
  });
}
