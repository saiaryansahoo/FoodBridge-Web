import API_BASE_URL from "./config.js";

document.getElementById("contactForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const name = document.getElementById("contactName").value.trim();
    const email = document.getElementById("contactEmail").value.trim();
    const message = document.getElementById("contactMessage").value.trim();

    if (!name || !email || !message) {
        alert("All fields are required!");
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/contact`, { // ❌ Removed trailing slash
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, email, message }),
        });

        if (!response.ok) throw new Error("Failed to send message");

        const data = await response.json();
        alert(data.message);
        document.getElementById("contactForm").reset(); // ✅ Clear form after success
    } catch (error) {
        console.error("Error:", error);
        alert("Could not send message. Please try again.");
    }
});
