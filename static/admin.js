// Submit Blog Post
document.getElementById('post-form').addEventListener('submit', async (e) => {
  e.preventDefault();
  const title = document.getElementById('post-title').value;
  const body = document.getElementById('post-body').value;

  const res = await fetch('/api/posts', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title, body })
  });

  if (res.ok) {
    alert('Post published!');
    e.target.reset();
  } else {
    alert('Failed to create post.');
  }
});

// Submit Item with Image
document.getElementById('item-form').addEventListener('submit', async (e) => {
  e.preventDefault();
  const name = document.getElementById('item-name').value;
  const image = document.getElementById('item-image').files[0];
  const formData = new FormData();

  formData.append('product_name', name);
  formData.append('image', image);

  const res = await fetch('/api/items', {
    method: 'POST',
    body: formData
  });

  if (res.ok) {
    alert('Item added!');
    e.target.reset();
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
async function deleteItem(id) {
  const confirmDelete = confirm("Are you sure you want to delete this item?");
  if (!confirmDelete) return;

  const res = await fetch(`/api/items/${id}`, {
    method: 'DELETE'
  });

  if (res.ok) {
    alert('Item deleted.');
    fetchItems();
  } else {
    alert('Failed to delete item.');
  }
}

// Load items on start
fetchItems();
