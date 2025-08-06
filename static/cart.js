document.addEventListener("DOMContentLoaded", () => {
    const cartButton = document.getElementById("cart");
    const cartCount = document.getElementById("cart-count");
    const cartItemsContainer = document.getElementById("cart-items");
    let cart = JSON.parse(localStorage.getItem("cart")) || [];
    
    function updateCartCount() {
        cartCount.textContent = cart.length;
    }

    function renderCartItems() {
        if (!cartItemsContainer) return;
        cartItemsContainer.innerHTML = "";
        let total = 0;
        cart.forEach((item, idx) => {
            total += item.price;
            const itemDiv = document.createElement("div");
            itemDiv.classList.add("cart-product");
            itemDiv.innerHTML = `
                <h3>${item.product_name}</h3>
                <p>${item.product_description}</p>
                <p>Price: $${(item.price / 100).toFixed(2)}</p>
                <p>${item.in_stock ? "In stock" : "Out of stock"}</p>
            `;
            const removeBtn = document.createElement("button");
            removeBtn.textContent = "Remove";
            removeBtn.className = "remove-cart-item";
            removeBtn.onclick = () => {
                cart.splice(idx, 1);
                localStorage.setItem("cart", JSON.stringify(cart));
                updateCartCount();
                renderCartItems();
            };
            itemDiv.appendChild(removeBtn);
            cartItemsContainer.appendChild(itemDiv);
        });
        const cartTotalElem = document.getElementById("cart-total");
        if (cartTotalElem) {
            cartTotalElem.textContent = `$${(total / 100).toFixed(2)}`;
        }
    }

    updateCartCount();
    renderCartItems();

    cartButton.addEventListener("click", () => {
        window.location.href = "/cart";
    });
});