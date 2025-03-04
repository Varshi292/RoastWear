const sqlite3 = require('sqlite3').verbose();
const readline = require('readline');
const path = require('path');

// Database file path (assuming it's in the same directory as this script)
const dbPath = path.join(__dirname, 'users.db');
const db = new sqlite3.Database(dbPath);

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

// Function to list all tables in the database
function listTables() {
    db.all("SELECT name FROM sqlite_master WHERE type='table'", [], (err, rows) => {
        if (err) {
            console.error("Error fetching tables:", err.message);
            process.exit(1);
        }
        
        if (rows.length === 0) {
            console.log("No tables found in the database.");
            process.exit(0);
        }

        console.log("Available tables:");
        rows.forEach((row, index) => {
            console.log(`${index + 1}. ${row.name}`);
        });
        
        // Prompt user to select a table
        rl.question("Enter the number of the table to delete: ", (answer) => {
            const index = parseInt(answer) - 1;
            if (isNaN(index) || index < 0 || index >= rows.length) {
                console.log("Invalid selection. Exiting.");
                process.exit(1);
            }
            
            const tableName = rows[index].name;
            deleteTable(tableName);
        });
    });
}

// Function to delete a table
function deleteTable(tableName) {
    db.run(`DROP TABLE IF EXISTS ${tableName}`, (err) => {
        if (err) {
            console.error(`Error deleting table ${tableName}:`, err.message);
        } else {
            console.log(`Table '${tableName}' deleted successfully.`);
        }
        rl.close();
        db.close();
    });
}

// Start the script
listTables();
