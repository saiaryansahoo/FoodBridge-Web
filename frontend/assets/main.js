const API_BASE_URL = "http://localhost:8080"; // Change this if your backend is hosted elsewhere

function authFetch(url, options = {}) {
    const token = localStorage.getItem("token"); // Get token from localStorage
    if (!token) {
        alert("You are not logged in!");
        window.location.href = "signup.html"; // Redirect to login page
        return;
    }

    // Add Authorization header
    options.headers = {
        ...options.headers,
        "Authorization": `Bearer ${token}`,
    };

    return fetch(`${API_BASE_URL}${url}`, options)
        .then(response => {
            if (response.status === 401) {
                alert("Unauthorized! Please log in again.");
                localStorage.removeItem("token");
                window.location.href = "signup.html";
                return;
            }
            return response.json();
        })
        .catch(error => console.error("Fetch Error:", error));
}
