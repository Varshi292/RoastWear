const sqlite3 = require('sqlite3').verbose();
const path = require('path');

// Path to the user database file
const dbPath = path.resolve(__dirname, './users.db'); // Points to ../Databases/user.db

// Open the database (will create it if it doesn't exist)
const db = new sqlite3.Database(dbPath, (err) => {
    if (err) {
        console.error('Error opening database:', err.message);
    } else {
        console.log('Connected to the user database.');
    }
});

// Function to initialize the database if needed
function initializeDatabase() {
    const createTableQuery = `
        CREATE TABLE IF NOT EXISTS user_uploads (
            username TEXT NOT NULL,
            filepath TEXT NOT NULL,
            PRIMARY KEY (username, filepath)
        );
    `;

    db.run(createTableQuery, (err) => {
        if (err) {
            console.error("Error creating table:", err.message);
        } else {
            console.log("Database initialized, user_uploads table is ready.");
        }
    });
}

// Call this function to initialize the database when the server starts
initializeDatabase();

// Export the database instance for use in other files
module.exports = db;
