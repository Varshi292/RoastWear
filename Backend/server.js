const express = require("express");
const bodyParser = require("body-parser");
const path = require("path");
const fs = require("fs");
const corsMiddleware = require("./middleware/corsConfig");
const routes = require("./Pages/index");

const app = express();
const port = 7777; // HTTP port
const host = "0.0.0.0"; // Listen on all network interfaces

// Middleware
app.use(corsMiddleware);
app.use(bodyParser.json());
app.use(express.static(path.join(__dirname, "../Frontend")));

// Use routes
app.use("/", routes);

// Start the server using HTTP
app.listen(port, host, () => {
    console.log(`Server running on http://${host}:${port}`);
});
