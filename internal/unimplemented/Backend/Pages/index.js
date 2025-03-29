const express = require("express");
const registerRoute = require("./registerManager");
const loginRoute = require("./loginManager");
const manageData = require("./userDataManager");
const manageUserMedia = require("./uploadMediaManager");
const path = require("path");

const router = express.Router();

router.use("/static", express.static(path.resolve(__dirname, "../Backend_Endpoints")));

router.get("/upload_image", (req, res) => {
    res.sendFile(path.resolve(__dirname, "../Backend_Endpoints/upload_media.html"));
});

// Handle registration API request
router.use("/register", registerRoute);

// Handle login API request
router.use("/login", loginRoute);

// Handle data management requests
router.use("/submitData", manageData);

// Serve static files from the "Backend/Backend_Endpoints" directory


// Serve the upload page at /upload_image


console.log("Serving static files from:", path.resolve(__dirname, "Backend/Backend_Endpoints"));

// Handle image upload requests
router.use("/upload_image", manageUserMedia);

module.exports = router;
