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
        CREATE TABLE IF NOT EXISTS user_data (
            username TEXT PRIMARY KEY,
            wake_up_time INTEGER,
            wake_up_confidence INTEGER,
            wake_up_intelligence INTEGER,
            wake_up_energy INTEGER,
            eat_breakfast_time INTEGER,
            eat_breakfast_confidence INTEGER,
            eat_breakfast_intelligence INTEGER,
            eat_breakfast_energy INTEGER,
            eat_lunch_time INTEGER,
            eat_lunch_confidence INTEGER,
            eat_lunch_intelligence INTEGER,
            eat_lunch_energy INTEGER,
            eat_dinner_time INTEGER,
            eat_dinner_confidence INTEGER,
            eat_dinner_intelligence INTEGER,
            eat_dinner_energy INTEGER,
            go_to_sleep_time INTEGER,
            go_to_sleep_confidence INTEGER,
            go_to_sleep_intelligence INTEGER,
            go_to_sleep_energy INTEGER,
            date TEXT
        );
    `;

    db.run(createTableQuery, (err) => {
        if (err) {
            console.error("Error creating table:", err.message);
        } else {
            console.log("Database initialized, user_data table is ready.");
        }
    });
}

// Call this function to initialize the database when the server starts
initializeDatabase();

// Export the database instance for use in other files
module.exports = db;
