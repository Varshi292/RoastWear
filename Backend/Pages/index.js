const express = require("express");
const registerRoute = require("./registerManager");
const loginRoute = require("./loginManager");
const manageData = require("./userDataManager");

const router = express.Router();

// Handle registration API request
router.use("/register", registerRoute);

// Handle login API request
router.use("/login", loginRoute);

// Handle data management requests
router.use("/submitData", manageData);

module.exports = router;
