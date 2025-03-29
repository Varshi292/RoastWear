const sqlite3 = require("sqlite3").verbose();

// Connect to the SQLite database
const db = new sqlite3.Database("users.db", (err) => {
    if (err) {
        console.error("Error connecting to the database:", err.message);
        process.exit(1); // Exit if connection fails
    }
});

// Function to delete a user by username
function deleteUser(username) {
    if (!username) {
        console.error("Error: Please provide a username to delete.");
        console.log("Usage: node deleteUser.js <username>");
        process.exit(1);
    }

    db.run("DELETE FROM users WHERE username = ?", [username], function (err) {
        if (err) {
            console.error("Error deleting user:", err.message);
        } else if (this.changes === 0) {
            console.log(`User '${username}' not found.`);
        } else {
            console.log(`✅ User '${username}' deleted successfully.`);
        }

        // Close database connection after operation
        db.close();
    });
}

// Get username from command-line arguments
const usernameToDelete = process.argv[2];

// Run the function
deleteUser(usernameToDelete);
