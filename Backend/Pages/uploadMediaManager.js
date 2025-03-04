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
        if (file.mimetype === 'image/png') {
            cb(null, true); // Accept PNG files
        } else {
            cb(null, false); // Reject non-PNG files
        }
    }
});



router.post("/", (req, res) => {
    upload.single('image')(req, res, (err) => {
        if (err) {
            return res.status(500).json({ error: 'An error occurred during file upload.' });
        }

        // If file was rejected by Multer, req.file will be undefined
        if (!req.file) {
            return res.status(500).json({ error: "Only PNG images are allowed." });
        }

        const { username } = req.body;
        if (!username) {
            return res.status(400).json({ error: "Username is required." });
        }

        const filename = req.file.filename; // Store only filename, not full path

        // Insert into database
        const insertQuery = `
            INSERT INTO user_uploads (username, filepath)
            VALUES (?, ?)
        `;
        db.run(insertQuery, [username, filename], (err) => {
            if (err) {
                console.error("Error inserting data into the database:", err.message);
                return res.status(500).json({ error: "Failed to store image data." });
            }
            res.status(200).json({ message: "Image uploaded and data saved.", filepath: filename });
        });
    });
});


module.exports = router;
