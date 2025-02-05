const express = require("express");
const bcrypt = require("bcrypt");
const db = require("../Databases/users");
const router = express.Router();

router.post("/", async (req, res) => {
  const { username, email, password } = req.body;

  // Input validation
  if (!username || !email || !password) {
    return res
      .status(400)
      .json({ success: false, message: "All fields are required." });
  }

  try {
    // Check if the email already exists
    db.get("SELECT * FROM users WHERE email = ?", [email], async (err, row) => {
      if (err) {
        console.error("Database error:", err);
        return res
          .status(500)
          .json({
            success: false,
            message: "Server error while checking email.",
          });
      }

      if (row) {
        return res
          .status(400)
          .json({ success: false, message: "Email already in use." });
      }

      // Hash the password and insert user into DB
      try {
        const hashedPassword = await bcrypt.hash(password, 10);
        db.run(
          "INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
          [username, email, hashedPassword],
          function (insertErr) {
            if (insertErr) {
              console.error("Insert error:", insertErr);
              return res
                .status(500)
                .json({ success: false, message: "Registration failed." });
            }
            res
              .status(201)
              .json({
                success: true,
                message: "User registered successfully.",
              });
          }
        );
      } catch (hashError) {
        console.error("Password hashing error:", hashError);
        res.status(500).json({ success: false, message: "Server error." });
      }
    });
  } catch (generalError) {
    console.error("General server error:", generalError);
    res
      .status(500)
      .json({ success: false, message: "Unexpected server error." });
  }
});

module.exports = router;
