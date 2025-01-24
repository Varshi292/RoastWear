const express = require("express");
const bodyParser = require("body-parser");
const path = require("path");
const fs = require("fs");
const https = require("https");
const corsMiddleware = require("./middleware/corsConfig");
const routes = require("./Pages/index");

const app = express();
const port = 3000; // HTTPS port
const host = "0.0.0.0"; // Listen on all network interfaces

// Path to your SSL certificate and key files
const privateKey = fs.readFileSync('/etc/letsencrypt/live/bgrillio.com/privkey.pem', 'utf8');
const certificate = fs.readFileSync('/etc/letsencrypt/live/bgrillio.com/fullchain.pem', 'utf8');
const ca = fs.readFileSync('/etc/letsencrypt/live/bgrillio.com/chain.pem', 'utf8'); // This is the CA cert, which is often included in the fullchain.pem

// SSL credentials
const credentials = { key: privateKey, cert: certificate, ca: ca };

// Middleware
app.use(corsMiddleware);
app.use(bodyParser.json());
app.use(express.static(path.join(__dirname, "../Frontend")));

// Use routes
app.use("/", routes);

// Start the server using HTTPS
https.createServer(credentials, app).listen(port, host, () => {
    console.log(`Server running on https://${host}:${port}`);
});
