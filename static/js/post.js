export function createPost() {
  const app = document.getElementById("app");
  app.innerHTML = `
      <h2>Créer un post</h2>
      <form id="postForm">
        <label for="title">Titre :</label><br>
        <input type="text" id="title" name="title" required><br><br>
  
        <label for="content">Contenu :</label><br>
        <textarea id="content" name="content" required></textarea><br><br>

        <label for="category">Catégorie :</label><br>
        <select name="category" id="category" required>
          <option disabled selected value="">--Sélectionner--</option>
          <option value="Actualité">Actualité</option>
          <option value="Sport">Sport</option>
          <option value="Culture">Culture</option>
          <option value="Technologie">Technologie</option>
          <option value="Autre">Autre</option>
        </select><br><br>
        <button type="submit">Créer le post</button>
      </form>
      <div id="postMessage"></div>
    `;
  const postForm = document.getElementById("postForm");
  postForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const post = {
      title: document.getElementById("title").value,
      content: document.getElementById("content").value,
      category: document.getElementById("category").value,
    };

    const res = await fetch("/post", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(post),
    });

    const message = document.getElementById("postMessage");

    if (res.ok) {
      message.innerText = "Post créé avec succès !";
      // Redirection vers la page d’accueil (SPA)
      navigateTo("home");
    } else {
      const err = await res.text();
      message.innerText = "Erreur : " + err;
    }
  });
}
