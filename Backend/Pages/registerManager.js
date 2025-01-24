const express = require("express");
const bcrypt = require("bcrypt");
const db = require("../Databases/users");

const router = express.Router();

router.post("/", async (req, res) => {
    const { username, email, password } = req.body;
    if (!username || !password || !email) {
        return res.status(400).json({ success: false, message: "All fields are required" });
    }

    db.get("SELECT * FROM users WHERE email = ?", [email], async (err, row) => {
        if (err) return res.status(500).json({ success: false, message: err.message });
        if (row) return res.status(400).json({ success: false, message: "Email already in use" });

        try {
            const hashedPassword = await bcrypt.hash(password, 10);
            db.run(
                "INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
                [username, email, hashedPassword],
                function (err) {
                    if (err) return res.status(500).json({ success: false, message: err.message });
                    res.status(201).json({ success: true, message: "User registered successfully" });
                }
            );
        } catch (hashErr) {
            res.status(500).json({ success: false, message: hashErr.message });
        }
    });
});

module.exports = router;
