const express = require('express');
const path = require('path');
const db = require('../Databases/user_data_table_init');

const router = express.Router();

// Directory where user images are stored
const imageDir = path.resolve(__dirname, '../user_images');

// GET endpoint to retrieve all uploaded images for a user
router.get("/", (req, res) => {
    const { username } = req.query;

    if (!username) {
        return res.status(400).json({ error: "Username is required." });
    }

    // Query the database for all uploads by the user
    const selectQuery = `
        SELECT filepath FROM user_uploads WHERE username = ?
    `;

    db.all(selectQuery, [username], (err, rows) => {
        if (err) {
            console.error("Error retrieving data from the database:", err.message);
            return res.status(500).json({ error: "Failed to retrieve image data." });
        }
        
        if (rows.length === 0) {
            return res.status(404).json({ message: "No images found for this user." });
        }
        
        res.status(200).json({ images: rows.map(row => row.filepath) });
    });
});

module.exports = router;
