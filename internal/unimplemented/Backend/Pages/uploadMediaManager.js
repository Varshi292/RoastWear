const express = require('express');
const multer = require('multer');
const path = require('path');
const fs = require('fs');
const db = require('../Databases/user_data_table_init'); // Adjust path to your database module

const router = express.Router();

// Directory to store uploaded images
const imageDir = path.resolve(__dirname, '../user_images');

// Ensure the directory exists
if (!fs.existsSync(imageDir)) {
    fs.mkdirSync(imageDir, { recursive: true }); // Create the directory if it doesn't exist
}

// Set up multer for handling file uploads
const storage = multer.diskStorage({
    destination: (req, file, cb) => {
        cb(null, imageDir); // Set destination to ../user_images
    },
    filename: (req, file, cb) => {
        const uniqueName = `${Date.now()}-${file.originalname}`; // Ensure unique filenames
        cb(null, uniqueName);
    },
});

const upload = multer({
    storage: storage,
    fileFilter: (req, file, cb) => {
        // Accept only PNG files
        if (file.mimetype === 'image/png') {
            cb(null, true);
        } else {
            cb(new Error('Only PNG images are allowed.'));
        }
    },
});

// POST endpoint to handle image uploads
router.post("/", upload.single('image'), (req, res) => {
    const { username } = req.body;
    const filepath = path.join(imageDir, req.file.filename);

    console.log("Recieved");

    if (!username) {
        return res.status(400).json({ error: "Username is required." });
    }

    // Insert the new record into the database
    const insertQuery = `
        INSERT INTO user_uploads (username, filepath)
        VALUES (?, ?)
    `;
    db.run(insertQuery, [username, filepath], (err) => {
        if (err) {
            console.error("Error inserting data into the database:", err.message);
            return res.status(500).json({ error: "Failed to store image data." });
        }
        res.status(200).json({ message: "Image uploaded and data saved.", filepath });
    });
});

module.exports = router;
