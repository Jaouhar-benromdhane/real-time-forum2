function logout() {
  fetch("/logout", {
    method: "POST",
    credentials: "include",
  })
    .then((res) => {
      if (res.ok) {
        // Redirection vers l'accueil après déconnexion
        navigateTo("home"); // si tu utilises une fonction SPA
        // OU
        // window.location.href = '/'; // si tu veux un vrai rechargement de page
      } else {
        console.error("Erreur lors de la déconnexion");
      }
    })
    .catch((err) => {
      console.error("Erreur réseau :", err);
    });
}
