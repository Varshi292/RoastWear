const sqlite3 = require('sqlite3').verbose();
const readline = require('readline');

const DB_PATH = './users.db'; // Change this to your database path

const db = new sqlite3.Database(DB_PATH, (err) => {
  if (err) {
    console.error('Failed to connect to database:', err.message);
    process.exit(1);
  }
});

function ask(question) {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });
  return new Promise(resolve => rl.question(question, ans => {
    rl.close();
    resolve(ans);
  }));
}

db.all(
  `SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%';`,
  async (err, tables) => {
    if (err) {
      console.error('Error reading tables:', err.message);
      process.exit(1);
    }

    if (tables.length === 0) {
      console.log('No tables found in the database.');
      process.exit(0);
    }

    console.log('Tables found:');
    tables.forEach((t, i) => console.log(`${i + 1}: ${t.name}`));

    const answer = await ask('Select a table by number: ');
    const index = parseInt(answer) - 1;

    if (isNaN(index) || index < 0 || index >= tables.length) {
      console.error('Invalid selection.');
      process.exit(1);
    }

    const selectedTable = tables[index].name;
    db.all(`SELECT * FROM ${selectedTable}`, (err, rows) => {
      if (err) {
        console.error('Error reading table contents:', err.message);
        process.exit(1);
      }

      console.log(`\nContents of "${selectedTable}":`);
      console.table(rows);
      db.close();
    });
  }
);
