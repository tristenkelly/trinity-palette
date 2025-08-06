// static/app.js
const itemsContainer = document.getElementById('gallery');

async function fetchAndDisplayItems() {
    try {
        const response = await fetch('/api/items'); // Fix: ensure correct path
        const items = await response.json();

        // Clear existing content
        itemsContainer.innerHTML = '';

        items.forEach(item => {
            const itemDiv = document.createElement('div');
            itemDiv.classList.add('product');

            const itemImage = document.createElement('img')
            itemImage.src = item.image_url;
            itemImage.alt = item.product_name || "Product image";

            const itemName = document.createElement('h3');
            itemName.textContent = item.product_name;

            const itemDescription = document.createElement('p');
            itemDescription.textContent = item.product_description;

            const itemPrice = document.createElement('p');
            itemPrice.textContent = `$${(item.price / 100).toFixed(2)}`; // assuming cents

            const itemStock = document.createElement('p');
            itemStock.textContent = item.in_stock ? "In stock" : "Out of stock";

            const addToCartButton = document.createElement('button');
            addToCartButton.textContent = "Add to Cart";
            addToCartButton.onclick = () => {
                const cart = JSON.parse(localStorage.getItem('cart')) || [];
                if (!cart.find(cartItem => cartItem.id === item.id)) {
                    cart.push(item);
                    localStorage.setItem('cart', JSON.stringify(cart));
                    alert(`${item.product_name} added to cart!`);
                } else {
                    alert(`${item.product_name} is already in your cart.`);
                }
            };

            itemDiv.appendChild(itemName);
            itemDiv.appendChild(itemDescription);
            itemDiv.appendChild(itemPrice);
            itemDiv.appendChild(itemStock);
            itemDiv.appendChild(itemImage);
            itemDiv.appendChild(addToCartButton);

            itemsContainer.appendChild(itemDiv);
        });

    } catch (error) {
        console.error('Error fetching or displaying items:', error);
    }
}

document.addEventListener('DOMContentLoaded', fetchAndDisplayItems);
