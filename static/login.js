const loginForm = document.getElementById('login-form');
document.addEventListener("DOMContentLoaded", () => {
  const loginForm = document.getElementById("login-form");

  if (!loginForm) {
    console.error("Login form not found.");
    return;
  }

  loginForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = document.getElementById('login-email').value;
    const password = document.getElementById('login-password').value;

    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });

    if (res.ok) {
      const data = await res.json();
      localStorage.setItem('token', data.token);
      localStorage.setItem('refresh_token', data.refresh_token);
      alert('Logged in successfully!');
      window.location.href = '/'; // or redirect to dashboard/homepage
    } else if (res.status === 401) {
      alert('Incorrect email or password.');
    } else {
      alert('Something went wrong during login.');
    }
});
  });