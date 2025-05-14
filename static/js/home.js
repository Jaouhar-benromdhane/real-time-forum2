import { InitComment } from "./comment.js";
import { chatBox } from "./message.js";

export async function loadHome() {
  try {
    let resp = await fetch("/home");

    // Si la réponse n'est pas OK, cela signifie probablement que l'utilisateur n'est pas connecté
    if (!resp.ok) {
      document.querySelector(".online").style.display = "none";
      document.querySelector(".offline").style.display = "block";

      // Afficher un message dans l'app
      const app = document.getElementById("app");
      app.innerHTML = "<p>Veuillez vous connecter pour voir le contenu.</p>";
      return;
    }

    let r = await resp.json();
    let posts = r.Posts;
    const app = document.getElementById("app");

    if (posts && posts.length > 0) {
      app.innerHTML = await formatPosts(posts);
      displayUsers();

      // Basculer l'affichage de navigation vers online
      document.querySelector(".online").style.display = "block";
      document.querySelector(".offline").style.display = "none";
      InitComment();
    } else {
      app.innerHTML = "<p>Aucun post disponible.</p>";

      // Si l'utilisateur est connecté mais qu'il n'y a pas de posts
      document.querySelector(".online").style.display = "block";
      document.querySelector(".offline").style.display = "none";
    }
  } catch (error) {
    console.error("Erreur lors du chargement de la page d'accueil:", error);

    // En cas d'erreur, afficher la navigation offline par défaut
    document.querySelector(".online").style.display = "none";
    document.querySelector(".offline").style.display = "block";

    const app = document.getElementById("app");
    app.innerHTML =
      "<p>Une erreur s'est produite lors du chargement des posts.</p>";
  }
}

async function formatPosts(posts) {
  let result = "";
  for (let i = 0; i < posts.length; i++) {
    let post = posts[i];
    let comments = await fetchComments(post.id);
    let postHTML = `
      <div class="post">
        <h1 class="title">${post.title}</h1>
        <h2 class="user">${post.user.nickname}</h2>
        <p class="content">${post.content}</p>
        <div class="footer">
          <span class="date">${post.date}</span>
          <span class="category">${post.category}</span>
        </div>

        <!-- Zone des commentaires -->
        <div class="comments" id="comments-${post.id}">
        ${formatComment(comments)}
        </div>

        <!-- Formulaire de commentaire -->
        <form class="comment-form" data-post-id="${post.id}">
          <input type="text" name="content" placeholder="Ajouter un commentaire" required />
          <button>Envoyer</button>
        </form>
      </div>
    `;
    result += postHTML;
  }
  return result;
}

async function fetchComments(postId) {
  let response = await fetch(`/comment/${postId}`);
  let r = await response.json();
  return r.comments ? r.comments : [];
}

function formatComment(comments) {
  let result = "";
  for (let i = 0; i < comments.length; i++) {
    let comment = comments[i];
    let commentHTML = `
      <div class="comment">
        <h1 class="user">${comment.user.nickname}</h1>
        <p class="content">${comment.content}</p>
        <span class="date">${comment.created_at}</span>
      </div>
    `;
    result += commentHTML;
  }
  return result;
}

function formatUsers(users) {
  // Trier les utilisateurs
  users.sort((a, b) => {
    // D'abord par statut de connexion (connectés en premier)
    if (a.connected !== b.connected) {
      return b.connected - a.connected;
    }
    // Ensuite par ordre alphabétique pour les non connectés
    if (!a.connected && !b.connected) {
      return a.user.nickname.localeCompare(b.user.nickname);
    }
    return 0;
  });

  // Déplacer l'expéditeur du dernier message en haut de la liste
  const lastMessageSender = localStorage.getItem("lastMessageSender");
  if (lastMessageSender) {
    const senderIndex = users.findIndex(
      (u) => u.user.nickname === lastMessageSender
    );
    if (senderIndex !== -1) {
      const [sender] = users.splice(senderIndex, 1);
      users.unshift(sender);
    }
  }

  let result = "";
  for (let i = 0; i < users.length; i++) {
    let user = users[i].user;
    let isConnected = users[i].connected;

    let userDiv = document.createElement("div");
    userDiv.className = "users_user";
    userDiv.onclick = () => chatBox(user.nickname);

    let nicknameHeading = document.createElement("h1");
    nicknameHeading.className = "nickname";
    nicknameHeading.setAttribute("data-nickname", user.nickname);
    nicknameHeading.textContent = user.nickname;

    let statusSpan = document.createElement("span");
    statusSpan.className = isConnected ? "connected" : "disconnected";
    statusSpan.textContent = isConnected ? "•" : "•";

    userDiv.appendChild(nicknameHeading);
    userDiv.appendChild(statusSpan);

    result += userDiv.outerHTML;
  }

  // Ajouter les listeners après l'insertion dans le DOM
  setTimeout(() => {
    const userElements = document.querySelectorAll(".users_user");
    userElements.forEach((element) => {
      element.addEventListener("click", () => {
        const nickname = element.querySelector(".nickname").textContent;
        chatBox(nickname);
      });
    });
  }, 0);

  return result;
}

export function displayUsers() {
  fetch("/refreshUsers")
    .then((response) => {
      if (!response.ok) {
        throw new Error("Erreur lors du chargement des utilisateurs");
      }
      return response.json();
    })
    .then((data) => {
      if (data.Users && Array.isArray(data.Users)) {
        const app = document.getElementById("users");
        app.innerHTML = formatUsers(data.Users);
      } else {
        console.warn("Format de données inattendu:", data);
      }
    })
    .catch((error) => {
      console.error("Erreur lors de la mise à jour des utilisateurs:", error);
      // Ne pas vider le conteneur en cas d'erreur pour éviter un flash d'interface
    });
}
