document.addEventListener('DOMContentLoaded', async () => {
  const authButton = document.getElementById("auth-button");
  const adminButton = document.getElementById("admin-button");
  const token = localStorage.getItem("token");
  if (!authButton) return;

  const setToLogin = () => {
    authButton.textContent = "Login";
    authButton.onclick = () => {
      window.location.href = "/login";
    };
  };

  if (token) {
    try {
      const verifyRes = await fetch("/api/verify", {
        headers: { Authorization: `Bearer ${token}` },
      });

      if (!verifyRes.ok) throw new Error("Invalid token");

      // Set logout behavior
      authButton.textContent = "Logout";
      authButton.onclick = () => {
        localStorage.removeItem("token");
        localStorage.removeItem("refresh_token");
        window.location.reload();
      };

      // Now fetch user info
      const userInfoRes = await fetch("/api/userInfo", {
        headers: { Authorization: `Bearer ${token}` },
      });
      if (userInfoRes.ok) {
        const user = await userInfoRes.json();
        if (user.is_admin && adminButton) {
          adminButton.style.display = "inline-block";
        }
      }

    } catch (err) {
      console.error("Auth check failed:", err);
      localStorage.removeItem("token");
      localStorage.removeItem("refresh_token");
      setToLogin();
    }
  } else {
    setToLogin();
  }
});
