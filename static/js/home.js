import { InitComment } from "./comment.js";

export async function loadHome() {
  let resp = await fetch("/home");
  let r = await resp.json();
  let posts = r.Posts;
  const app = document.getElementById("app");
  if (r.Posts) {
    app.innerHTML = await formatPosts(posts);
    displayUsers(); // maybe to delete
    document.querySelector(".online").style.display = "block";
    document.querySelector(".offline").style.display = "none";
    InitComment()
  } else {
    document.querySelector(".online").style.display = "none";
    document.querySelector(".offline").style.display = "block";
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
  let result = "";
  for (let i = 0; i < users.length; i++) {
    let user = users[i].user;
    let isConnected = users[i].connected;
    let userHTML = `
      <div class="users_user">
        <h1 class="nickname">${user.nickname}</h1> ${
      isConnected
        ? `<span class="connected">•</span>`
        : `<span class="disconnected">•</span>`
    }
      </div>
    `;
    result += userHTML;
  }
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