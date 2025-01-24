// db.js
const sqlite3 = require('sqlite3').verbose();
const path = require('path');

// Path to your database file (resolved from the current script's directory)
const dbPath = path.resolve(__dirname, 'users.db');  // This will create it in the same directory as this script

// Open the database (will create it if it doesn't exist)
const db = new sqlite3.Database(dbPath, (err) => {
    if (err) {
        console.error('Error opening database:', err.message);
    } else {
        console.log('Connected to the database.');
    }
});

// Function to initialize the database if needed
function initializeDatabase() {
    const createTableQuery = `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `;

    db.run(createTableQuery, (err) => {
        if (err) {
            console.error("Error creating table:", err.message);
        } else {
            console.log("Database initialized, users table is ready.");
        }
    });
}

// Call this function to initialize the database when the server starts
initializeDatabase();

// Export the database instance for use in other files
module.exports = db;
