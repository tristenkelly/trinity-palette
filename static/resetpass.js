document.addEventListener("DOMContentLoaded", () => {
  const resetForm = document.getElementById("reset-form");

  if (!resetForm) {
    console.error("Reset form not found.");
    return;
  }

  resetForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = document.getElementById('reset-email').value;
    const password = document.getElementById('reset-password').value;

    const res = await fetch('/api/changepassword', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });

if (res.ok) {
  alert('Password changed successfully!');
  window.location.href = '/login'; // or wherever
} else {
  const text = await res.text(); // helpful for debugging
  console.error("Change password failed:", text);
  alert('Failed to change password.');
}

});
  });