export function InitComment() {
  document.querySelectorAll(".comment-form").forEach((form) => {
    form.addEventListener("submit", async (e) => {
      e.preventDefault();

      const post_id = e.target.attributes["data-post-id"].value;
      const content = e.target.content.value;
      let resp = await fetch("/comment", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ post_id: Number(post_id), content: content }),
      });
      let r = await resp.json();
      if (!r.error) {
        navigateTo("home");
      }
    });
  });
}
