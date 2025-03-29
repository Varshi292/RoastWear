const express = require("express");
const bcrypt = require("bcrypt");
const db = require("../Databases/users");

const router = express.Router();

router.post("/", (req, res) => {
  const { username, password } = req.body;
  if (!username || !password) {
    return res
      .status(400)
      .json({ success: false, message: "All fields are required." });
  }

  db.get(
    "SELECT * FROM users WHERE username = ?",
    [username],
    async (err, row) => {
      if (err) {
        return res
          .status(500)
          .json({ success: false, message: "Server error." });
      }

      if (!row) {
        return res
          .status(401)
          .json({ success: false, message: "Invalid username or password." });
      }

      try {
        const match = await bcrypt.compare(password, row.password);
        if (match) {
          res.status(200).json({ success: true, message: "Login successful." });
        } else {
          res
            .status(401)
            .json({ success: false, message: "Invalid username or password." });
        }
      } catch (error) {
        res.status(500).json({ success: false, message: "Server error." });
      }
    }
  );
});

module.exports = router;
