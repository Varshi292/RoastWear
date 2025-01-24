document.addEventListener("DOMContentLoaded", () => {
    const registerForm = document.getElementById("registerForm");

    registerForm.addEventListener("submit", async (e) => {
        e.preventDefault();

        const username = document.getElementById("username").value.trim();
        const email = document.getElementById("email").value.trim();
        const password = document.getElementById("password").value.trim();
        const messageBox = document.getElementById("message");

        if (!username || !email || !password) {
            showMessage("All fields are required!", "error");
            return;
        }

        try {
            const response = await fetch("/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username, email, password })
            });

            const result = await response.json();
            showMessage(result.message, result.success ? "success" : "error");

            if (result.success) {
                setTimeout(() => window.location.href = "/login", 2000);
            }
        } catch (error) {
            showMessage("An error occurred. Please try again.", "error");
            console.error(error);
        }
    });

    function showMessage(message, type) {
        messageBox.textContent = message;
        messageBox.className = type;
        messageBox.style.display = "block";
    }
});
