const sqlite3 = require('sqlite3').verbose();
const readline = require('readline');

// Predefined database path
const dbPath = './users.db';  // Replace with your actual database path

// Initialize the readline interface
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

// Open the database
const db = new sqlite3.Database(dbPath, sqlite3.OPEN_READONLY, (err) => {
    if (err) {
        console.error('Error opening database:', err.message);
    } else {
        console.log('Connected to the database');
        listTables();
    }
});

// Function to list all tables in the database
function listTables() {
    const sql = "SELECT name FROM sqlite_master WHERE type='table'";

    db.all(sql, [], (err, rows) => {
        if (err) {
            throw err;
        }

        console.log("Tables in the database:");
        rows.forEach((row, index) => {
            console.log(`${index + 1}. ${row.name}`);
        });

        // Prompt the user to choose a table
        rl.question('Please select a table by entering its number: ', (answer) => {
            const tableIndex = parseInt(answer) - 1;
            const tableName = rows[tableIndex] ? rows[tableIndex].name : null;

            if (tableName) {
                showTableData(tableName);
            } else {
                console.log('Invalid selection. Exiting...');
                rl.close();
            }
        });
    });
}

// Function to display data from the chosen table
function showTableData(tableName) {
    const sql = `SELECT * FROM ${tableName}`;

    db.all(sql, [], (err, rows) => {
        if (err) {
            throw err;
        }

        console.log(`Data from table: ${tableName}`);
        console.table(rows);  // Display the table data in a tabular format

        // Close the readline interface after displaying the data
        rl.close();
    });
}

// Close the database when the script finishes
db.on('close', () => {
    console.log('Closed the database connection');
});
