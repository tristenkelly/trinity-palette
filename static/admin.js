document.addEventListener('DOMContentLoaded', async () => {
  const token = localStorage.getItem('token');
  if (!token) {
    alert('You must be logged in to access admin features.');
    window.location.href = '/login';
    return;
  }

  try {
    const res = await fetch('/api/userInfo', {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });

    if (!res.ok) throw new Error('Not authorized');

    const user = await res.json();

    if (user.is_admin !== true) {
      alert('You do not have permission to access this page.');
      window.location.href = '/';
      return;
    }

    // Submit Blog Post
    document.getElementById('post-form').addEventListener('submit', async (e) => {
  e.preventDefault();
  const title = document.getElementById('post-title').value;
  const body = document.getElementById('post-body').value;
  const token = localStorage.getItem('token'); // Get the token

  const res = await fetch('/api/post', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`  // Attach the token here
    },
    body: JSON.stringify({ title, body })
  });

  // Optionally handle the response
  if (!res.ok) {
    const error = await res.text();
    console.error('Post failed:', error);
    alert('Failed to submit post');
  } else {
    alert('Post submitted!');
  }
});


    // Submit Item with Image
    document.getElementById('item-form').addEventListener('submit', async (e) => {
      e.preventDefault();
      const name = document.getElementById('item-name').value;
      const image = document.getElementById('item-image').files[0];
      const description = document.getElementById('item-description').value
      const price = document.getElementById('item-price').value
      const instock = document.getElementById('item-instock').value
      const formData = new FormData();

      formData.append('product_name', name);
      formData.append('product_description', description)
      formData.append('image', image);
      formData.append('price', price);
      formData.append('in_stock', instock);

      const res = await fetch('/admin/item/create', {
        method: 'POST',
        body: formData
      });

      if (res.ok) {
        alert('Item added!');
        e.target.reset();
        fetchItems(); // Refresh the list after adding
      } else {
        alert('Failed to add item.');
      }
    });

    // Fetch items for deletion
    async function fetchItems() {
      const res = await fetch('/api/items');
      const items = await res.json();
      const list = document.getElementById('item-list');
      list.innerHTML = '';

      items.forEach(item => {
        const div = document.createElement('div');
        div.classList.add('item');
        div.innerHTML = `
          <span>${item.product_name}</span>
          <button onclick="deleteItem('${item.id}')">Delete</button>
        `;
        list.appendChild(div);
      });
    }

    // Delete item
    window.deleteItem = async function (id) {
      const confirmDelete = confirm("Are you sure you want to delete this item?");
      if (!confirmDelete) return;

      const res = await fetch(`/api/item/${id}`, {
        method: 'DELETE'
      });

      if (res.ok) {
        alert('Item deleted.');
        fetchItems();
      } else {
        alert('Failed to delete item.');
      }
    };

    fetchItems();

  } catch (err) {
    console.error('Auth check failed:', err);
    window.location.href = '/login';
  }
});
