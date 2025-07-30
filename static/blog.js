document.addEventListener("DOMContentLoaded", () => {
  const blogContainer = document.getElementById("blog-posts");

  async function fetchPosts() {
    try {
      const res = await fetch("/api/posts");
      if (!res.ok) throw new Error("Failed to fetch blog posts");

      const posts = await res.json();
      posts.forEach(renderPost);
    } catch (err) {
      blogContainer.innerHTML = `<p style="color: red;">Oops! Could not load blog posts 💔</p>`;
      console.error(err);
    }
  }

  function renderPost(post) {
    const postElement = document.createElement("div");
    postElement.classList.add("blog-post");

    const createdDate = new Date(post.created_at).toLocaleDateString("en-US", {
      month: "long",
      day: "numeric",
      year: "numeric"
    });

    postElement.innerHTML = `
      <h2>${sanitize(post.title)}</h2>
      <p class="blog-date">By <strong>${sanitize(post.username)}</strong> on ${createdDate}</p>
      <p class="blog-content">${truncate(sanitize(post.body), 240)}</p>
      <a href="/blog/post?title=${encodeURIComponent(post.title)}" class="blog-link">Read more →</a>
    `;

    blogContainer.appendChild(postElement);
  }

  // Prevent XSS if data isn't sanitized server-side
  function sanitize(str) {
    const div = document.createElement("div");
    div.textContent = str;
    return div.innerHTML;
  }

  function truncate(str, maxLength) {
    return str.length > maxLength ? str.slice(0, maxLength) + "..." : str;
  }

  fetchPosts();
});
