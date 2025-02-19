import API_BASE_URL from "./config.js";

// ✅ Function to check authentication status
function checkAuth() {
    const token = localStorage.getItem("token");
    const dashboard = document.getElementById("dashboard");

    if (!token) {
        alert("You need to log in first!");
        window.location.href = "signup.html"; // Redirect to login page
    } else {
        dashboard.style.display = "block"; // Show dashboard if logged in
        fetchFoodItems(); // Fetch food items only after authentication
    }
}

// ✅ Function to fetch food items
async function fetchFoodItems() {
    const token = localStorage.getItem("token");

    try {
        const response = await fetch(`${API_BASE_URL}/food/`, {
            headers: { "Authorization": `Bearer ${token}` },
        });

        if (!response.ok) throw new Error("Failed to fetch food items");

        const foodItems = await response.json();
        displayFoodItems(foodItems);
    } catch (error) {
        console.error("Error fetching food:", error);
        alert("Could not load food listings. Please try again later.");
    }
}

// ✅ Function to display food items in the UI
function displayFoodItems(foodItems) {
    const foodList = document.getElementById("foodList");
    foodList.innerHTML = "";

    if (foodItems.length === 0) {
        foodList.innerHTML = "<p>No food items available.</p>";
        return;
    }

    foodItems.forEach((food) => {
        foodList.innerHTML += `
            <div class="food-Form">
                <h3>${food.name}</h3>
                <p>${food.description}</p>
                <p><strong>Quantity:</strong> ${food.quantity}</p>
                <p><strong>Location:</strong> ${food.location}</p>
            </div>`;
    });
}

// ✅ Function to handle food listing form submission
document.getElementById("foodForm")?.addEventListener("submit", async (e) => {
    e.preventDefault();

    const token = localStorage.getItem("token");
    if (!token) {
        alert("You need to log in first!");
        return;
    }

    const name = document.getElementById("foodName").value.trim();
    const description = document.getElementById("foodDescription").value.trim();
    const quantity = parseInt(document.getElementById("foodQuantity").value.trim(), 10);
    const location = document.getElementById("foodLocation").value.trim();


    if (!name || !description || !quantity || !location) {
        alert("All fields are required!");
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/food/list`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            body: JSON.stringify({ name, description, quantity, location }),
        });

        const data = await response.json();
        if (!response.ok) throw new Error(data.error || "Food listing failed");

        alert(data.message);
        document.getElementById("foodForm").reset();
        fetchFoodItems(); // Refresh the list
    } catch (error) {
        console.error("Error listing food:", error);
        alert("Failed to list food. Please try again.");
    }
});

// ✅ Ensure user is logged in before displaying the dashboard
document.addEventListener("DOMContentLoaded", checkAuth);
