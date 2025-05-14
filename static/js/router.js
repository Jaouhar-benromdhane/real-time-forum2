import { loadHome } from "./home.js";
import { loadLogin } from "./login.js";
import { loadRegister } from "./register.js";
import { createPost } from "./post.js";
import { ws } from "./websocket.js";

window.logout = async function () {
  const res = await fetch("/logout", { method: "POST" });
  if (res.ok) {
    // Switch la navigation vers offline avant de naviguer vers login
    document.querySelector(".online").style.display = "none";
    document.querySelector(".offline").style.display = "block";

    navigateTo("login");

    // Si WebSocket est initialisé, envoyer un message de déconnexion
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: "logout" }));
    }
  } else {
    alert("Erreur lors de la déconnexion");
  }
};

window.navigateTo = function (page) {
  document.getElementById("users").innerHTML = "";
  switch (page) {
    case "home":
      loadHome();
      break;
    case "login":
      loadLogin();
      // S'assurer que nous sommes bien en mode offline quand on navigue vers login
      document.querySelector(".online").style.display = "none";
      document.querySelector(".offline").style.display = "block";
      break;
    case "register":
      loadRegister();
      // S'assurer que nous sommes bien en mode offline quand on navigue vers register
      document.querySelector(".online").style.display = "none";
      document.querySelector(".offline").style.display = "block";
      break;
    case "post":
      createPost();
      break;
    default:
      document.getElementById("app").innerHTML = "<p>Page introuvable.</p>";
  }
};

// On charge une page par défaut au démarrage
window.addEventListener("DOMContentLoaded", () => {
  // Vérifier si l'utilisateur est connecté en essayant de charger la page home
  fetch("/home")
    .then((response) => {
      if (response.ok) {
        // Si connecté, afficher la nav online
        document.querySelector(".online").style.display = "block";
        document.querySelector(".offline").style.display = "none";
      } else {
        // Si non connecté, afficher la nav offline
        document.querySelector(".online").style.display = "none";
        document.querySelector(".offline").style.display = "block";
      }
      navigateTo("home");
    })
    .catch(() => {
      // En cas d'erreur, afficher la nav offline par défaut
      document.querySelector(".online").style.display = "none";
      document.querySelector(".offline").style.display = "block";
      navigateTo("login");
    });
});
