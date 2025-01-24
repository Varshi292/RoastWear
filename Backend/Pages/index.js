const path = require("path");
const express = require("express");
const registerRoute = require("./registerManager");
const loginRoute = require("./loginManager");
const manageData = require("./userDataManager.js");

const router = express.Router();

// Load the register page
router.get("/register", (req, res) => {
    res.sendFile(path.join(__dirname, "../../Frontend/registerpage.html"));
});

// Handle registration API request
router.use("/register", registerRoute);

//manage the login page
router.get("/login", (req, res) => {
    res.sendFile(path.join(__dirname, "../../Frontend/loginpage.html"));
});

router.use("/login", loginRoute);

//manage the home page
router.get("/", (req, res) => {
    res.sendFile(path.join(__dirname, "../../Frontend/homepage.html"));
});

// In your route manager (e.g., `dashboardManager.js`)
router.get("/dashboard", (req, res) => {
    res.sendFile(path.join(__dirname, "../../Frontend/dashboard.html"));
});

router.use(manageData);  // This will use the routes defined in userDataManager.js without adding another /submitData prefix


module.exports = router;