export function createComment(postId) {
  const app = document.getElementById("app");
  app.innerHTML = `
      <h2>Créer un commentaire</h2>
      <form id="commentForm">
        <label for="content">Contenu :</label><br>
        <textarea id="content" name="content" required></textarea><br><br>
        <button type="submit">Créer le commentaire</button>
      </form>
      <div id="commentMessage"></div>
    `;
  const commentForm = document.getElementById("commentForm");
  commentForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const comment = {
      content: document.getElementById("content").value,
    };

    const res = await fetch(`/post/${postId}/comment`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(comment),
    });

    const message = document.getElementById("commentMessage");

    if (res.ok) {
      message.innerText = "Commentaire créé avec succès !";
      // Redirection vers la page d’accueil (SPA)
      navigateTo(`post/${postId}`);
    } else {
      const err = await res.text();
      message.innerText = "Erreur : " + err;
    }
  });
}