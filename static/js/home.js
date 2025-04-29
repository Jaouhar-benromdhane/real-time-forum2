// export function loadHome() {
//   fetch("/home")
//     .then((response) => {
//       if (!response.ok) throw new Error("Erreur lors du chargement des posts");
//       console.log(response.json())
//       return response.json(); // car /home renvoie du HTML
//     })
//     .then((posts) => {
//       console.log("HTML reçu :", html);
//       const app = document.getElementById("app");
//       app.innerHTML = formatPosts(posts);

//     })
//     .catch((error) => {
//       console.error("Erreur :", error);
//       document.getElementById("app").innerHTML =
//         "<p>Impossible de charger les posts.</p>";
//     });
// }

export async function loadHome() {
  let resp = await fetch("/home");
  let r = await resp.json();
  let posts = r.Posts;
  const app = document.getElementById("app");
  if (r.Posts) {
    app.innerHTML = formatPosts(posts);
    displayUsers(); // maybe to delete
    document.querySelector(".online").style.display = "block";
    document.querySelector(".offline").style.display = "none";
  } else {
    document.querySelector(".online").style.display = "none";
    document.querySelector(".offline").style.display = "block";
  }
}

function formatPosts(posts) {
  let result = "";
  for (let i = 0; i < posts.length; i++) {
    let post = posts[i];
    let postHTML = `
      <div class="post">
      <h1 class="title">${post.title}</h1>
      <h2 class="user">${post.user.nickname}</h2>
    
      <p class="content">${post.content}</p>

      <div class="footer">
        <span class="date">${post.date}</span>
        <span class="category">${post.category}</span>
      </div>
    </div>
    `;
    result += postHTML;
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
