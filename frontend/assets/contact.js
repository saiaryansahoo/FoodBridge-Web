import API_BASE_URL from "./config.js";

document.addEventListener("DOMContentLoaded", () => {
    const contactForm = document.getElementById("contact-form");

    if (!contactForm) {
        console.error("❌ contact-form element not found!");
        return;
    }

    contactForm.addEventListener("submit", async (e) => {
        e.preventDefault();

        const name = document.querySelector("input[name='name']").value.trim();
        const email = document.querySelector("input[name='email']").value.trim();
        const message = document.querySelector("textarea[name='message']").value.trim();

        if (!name || !email || !message) {
            alert("⚠️ All fields are required!");
            return;
        }

        console.log("📩 Submitting contact form:", { name, email, message });

        try {
            const response = await fetch(`${API_BASE_URL}/contact`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ name, email, message }),
            });

            const data = await response.json();
            console.log("✅ Contact API Response:", data);

            if (!response.ok) throw new Error(data.error || "Contact submission failed");

            alert("🎉 Message sent successfully!");
            contactForm.reset();
        } catch (error) {
            console.error("❌ Error submitting contact form:", error);
            alert("🚨 Failed to send message. Please try again.");
        }
    });
});
