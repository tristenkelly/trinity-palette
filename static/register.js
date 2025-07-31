document.addEventListener('DOMContentLoaded', () => {
  const registerForm = document.getElementById('register-form');

  if (!registerForm) {
    console.error('Register form not found.');
    return;
  }

  registerForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = document.getElementById('register-email').value;
    const username = document.getElementById('register-username').value;
    const password = document.getElementById('register-password').value;

    const res = await fetch('/api/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, username, password })
    });

    if (res.ok) {
      const loginRes = await fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
      });

      if (loginRes.ok) {
        window.location.href = '/';
      } else {
        alert('Registration succeeded but automatic login failed. Please log in manually.');
        window.location.href = '/login';
      }
    } else {
      alert('Registration failed. Please try again.');
    }
  });
});
