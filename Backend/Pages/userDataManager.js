const express = require("express");
const db = require("../Databases/user_data_table_init"); // Import the db module
const router = express.Router();

// Route for submitting data
router.post("/submitData", async (req, res) => {
  const { username, selectedEvent, time, confidence, intelligence, energy } =
    req.body;
  const date = new Date().toISOString().split("T")[0]; // Get current date in YYYY-MM-DD format

  console.log("Received data:", req.body); // Log the data to check what is being sent

  // Check if the user already has data for the given date
  const checkQuery = `SELECT * FROM user_data WHERE username = ? AND date = ?`;
  db.get(checkQuery, [username, date], (err, row) => {
    if (err) {
      console.error("Error checking data:", err.message);
      return res
        .status(500)
        .json({ success: false, message: "Database error" });
    }

    let query, values;
    if (row) {
      // If data exists, update it
      console.log("Data exists, updating...");
      query = `UPDATE user_data SET `;
      values = [];

      // Determine the correct columns based on the selected event
      if (selectedEvent === "wake_up") {
        query += `wake_up_time = ?, wake_up_confidence = ?, wake_up_intelligence = ?, wake_up_energy = ?`;
        values = [time, confidence, intelligence, energy];
      } else if (selectedEvent === "eat_breakfast") {
        query += `eat_breakfast_time = ?, eat_breakfast_confidence = ?, eat_breakfast_intelligence = ?, eat_breakfast_energy = ?`;
        values = [time, confidence, intelligence, energy];
      } else if (selectedEvent === "eat_lunch") {
        query += `eat_lunch_time = ?, eat_lunch_confidence = ?, eat_lunch_intelligence = ?, eat_lunch_energy = ?`;
        values = [time, confidence, intelligence, energy];
      } else if (selectedEvent === "eat_dinner") {
        query += `eat_dinner_time = ?, eat_dinner_confidence = ?, eat_dinner_intelligence = ?, eat_dinner_energy = ?`;
        values = [time, confidence, intelligence, energy];
      } else if (selectedEvent === "go_to_sleep") {
        query += `go_to_sleep_time = ?, go_to_sleep_confidence = ?, go_to_sleep_intelligence = ?, go_to_sleep_energy = ?`;
        values = [time, confidence, intelligence, energy];
      }

      // Add date and username to the query
      query += `, date = ? WHERE username = ? AND date = ?`;
      values.push(date, username, date);
    } else {
      // If no data exists, insert a new record
      console.log("No data, inserting...");
      query = `INSERT INTO user_data (username, ${selectedEvent}_time, ${selectedEvent}_confidence, ${selectedEvent}_intelligence, ${selectedEvent}_energy, date) VALUES (?, ?, ?, ?, ?, ?)`;
      values = [username, time, confidence, intelligence, energy, date];
    }

    // Execute the query
    db.run(query, values, function (err) {
      if (err) {
        console.error("Error executing query:", err.message);
        return res
          .status(500)
          .json({ success: false, message: "Error saving data" });
      }

      res
        .status(200)
        .json({
          success: true,
          message: `${selectedEvent} data logged successfully`,
        });
    });
  });
});

module.exports = router;
