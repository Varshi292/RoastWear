const express = require("express");
const bcrypt = require("bcrypt");
const db = require("../Databases/users");

const router = express.Router();

router.post("/", async (req, res) => {
    const { username, password } = req.body;
    if (!username || !password) {
        return res.status(400).json({ success: false, message: "All fields are required" });
    }

    

    db.get("SELECT * FROM users WHERE username = ?", [username], async (err, row) => {
        if (err) {
            return res.status(500).json({ success: false, message: "Internal server error." });
        }
    
        if (!row) {
            return res.status(401).json({ success: false, message: "Email or password is incorrect." });
        }
    
        try {
            // Compare input password with hashed password from DB
            const match = await bcrypt.compare(password, row.password);
    
            if (match) {
                res.status(200).json({ success: true, message: "Login successful!" });
            } else {
                res.status(401).json({ success: false, message: "Email or password is incorrect." });
            }
        } catch (hashErr) {
            console.error("Error comparing passwords:", hashErr.message);
            res.status(500).json({ success: false, message: "Internal server error." });
        }
    });
    
});

module.exports = router;