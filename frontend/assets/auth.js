import API_BASE_URL from "./config.js";

document.addEventListener("DOMContentLoaded", () => {
    // ✅ Signup Form
    const signupForm = document.getElementById("register-form");
    if (signupForm) {
        signupForm.addEventListener("submit", async (e) => {
            e.preventDefault();

            const name = document.getElementById("registerName").value.trim();
            const email = document.getElementById("registerEmail").value.trim();
            const password = document.getElementById("registerPassword").value.trim();

            if (!name || !email || !password) {
                alert("All fields are required!");
                return;
            }

            try {
                const response = await fetch(`${API_BASE_URL}/auth/register`, {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ name, email, password }),
                });

                const data = await response.json();
                
                if (response.ok) {
                    localStorage.setItem("token", data.token); // ✅ Save token for authentication
                    alert("Signup successful!");
                    window.location.href = "dashboard.html"; // ✅ Redirect to dashboard
                } else {
                    alert(data.error || "Signup failed. Please try again.");
                }
            } catch (error) {
                console.error("Signup error:", error);
                alert("Signup failed. Please check your network and try again.");
            }
        });
    }

    // ✅ Login Form
    const loginForm = document.getElementById("loginForm");
    if (loginForm) {
        loginForm.addEventListener("submit", async (e) => {
            e.preventDefault();

            const email = document.getElementById("loginEmail").value.trim();
            const password = document.getElementById("loginPassword").value.trim();

            if (!email || !password) {
                alert("Both fields are required!");
                return;
            }

            try {
                const response = await fetch(`${API_BASE_URL}/auth/login`, {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ email, password }),
                });

                const data = await response.json();

                if (response.ok && data.token) {
                    localStorage.setItem("token", data.token); // ✅ Save token for authentication
                    alert("Login successful!");
                    window.location.href = "dashboard.html"; // ✅ Redirect to dashboard
                } else {
                    alert(data.error || "Login failed. Please check your credentials.");
                }
            } catch (error) {
                console.error("Login error:", error);
                alert("Login failed. Please check your network and try again.");
            }
        });
    }
    
    // ✅ Check if User is Logged In (Auto-Redirect)
    const token = localStorage.getItem("token");
    if (token && window.location.pathname.includes("login.html")) {
        window.location.href = "dashboard.html"; // ✅ If already logged in, redirect
    }
});
