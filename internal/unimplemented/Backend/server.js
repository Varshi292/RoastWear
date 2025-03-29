const express = require("express");
const bodyParser = require("body-parser");
const path = require("path");
const cors = require("cors");
const fs = require("fs");
const corsMiddleware = require("./middleware/corsConfig");
const routes = require("./Pages/index");
const loginRoute = require("./Pages/loginManager");
const registerRoute = require("./Pages/registerManager");
const host = "localhost";

const app = express();
const port = 7777; // HTTP port
//const host = "0.0.0.0";  Listen on all network interfaces

// Middleware
app.use(cors({ origin: "http://localhost:3000" }));
app.use(corsMiddleware);
app.use(bodyParser.json());
app.get("/", (req, res) => {
  res.send("Backend Homepage!");
});

app.use(routes);

// Start the server using HTTP
app.listen(port, host, () => {
  console.log(`Server running on http://localhost:${port}`);
});