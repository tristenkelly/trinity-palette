document.addEventListener('DOMContentLoaded', async () => {
  const authButton = document.getElementById("auth-button");
  const token = localStorage.getItem("token");

  if (!authButton) return;

  if (token) {
    try {
      const res = await fetch("/api/verify", {
        headers: { Authorization: `Bearer ${token}` },
      });

      if (res.ok) {
        authButton.textContent = "Logout";
        authButton.onclick = () => {
          localStorage.removeItem("token");
          localStorage.removeItem("refresh_token");
          window.location.reload();
        };
      } else {
        throw new Error("Invalid token");
      }
    } catch {
      localStorage.removeItem("token");
      localStorage.removeItem("refresh_token");
      authButton.textContent = "Login";
      authButton.onclick = () => {
        window.location.href = "/login";
      };
    }
  } else {
    authButton.textContent = "Login";
    authButton.onclick = () => {
      window.location.href = "/login";
    };
  }
});
