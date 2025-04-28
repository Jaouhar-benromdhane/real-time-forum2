export function loadRegister() {
  const app = document.getElementById("app");
  app.innerHTML = `
      <h2>Inscription</h2>
      <form id="registerForm">
        <label>Pseudo :</label><br>
        <input type="text" id="nickname" required><br>
  
        <label>Prénom :</label><br>
        <input type="text" id="firstName" required><br>
  
        <label>Nom :</label><br>
        <input type="text" id="lastName" required><br>
  
        <label>Âge :</label><br>
        <input type="number" id="age" required><br>
  
        <label>Genre :</label><br>
         <select name="gender" id= "gender" required>
      <option disabled selected value="">--Sélectionner--</option>
      <option value="Male">Homme</option>
      <option value="Female">Femme</option>
      <option value="Other">Autre</option>
    </select><br>
  
        <label>Email :</label><br>
        <input type="email" id="email" required><br>
  
        <label>Mot de passe :</label><br>
        <input type="password" id="password" required><br><br>
  
        <button type="submit">S'inscrire</button>
      </form>
      <div id="registerMessage"></div>
    `;

  const form = document.getElementById("registerForm");
  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const user = {
      nickname: document.getElementById("nickname").value,
      first_name: document.getElementById("firstName").value,
      last_name: document.getElementById("lastName").value,
      age: Number(document.getElementById("age").value),
      gender: document.getElementById("gender").value,
      email: document.getElementById("email").value,
      password: document.getElementById("password").value,
    };

    const res = await fetch("/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    });

    const message = document.getElementById("registerMessage");

    if (res.ok) {
      message.innerText = "Inscription réussie !";
      navigateTo("login");
    } else {
      const err = await res.text();
      message.innerText = "Erreur : " + err;
    }
  });
}
