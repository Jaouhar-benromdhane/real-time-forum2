import { loadHome } from "./home.js";
import { loadLogin } from "./login.js";
import { loadRegister } from "./register.js";
import { createPost } from "./post.js";
import { ws } from "./websocket.js";

window.logout = async function () {
  const res = await fetch("/logout", { method: "POST" });
  if (res.ok) {
    navigateTo("login");
    ws.send(JSON.stringify({ type: "logout" }));
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
      break;
    case "register":
      loadRegister();
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
  navigateTo("home");
});
