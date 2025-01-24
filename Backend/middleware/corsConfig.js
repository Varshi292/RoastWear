const cors = require("cors");

const corsOptions = {
    origin: "*", // Allows requests from any origin (*). Change to specific domain in production.
    methods: "GET, POST, PUT, DELETE", // Allowed request methods
    allowedHeaders: "Content-Type, Authorization", // Allowed headers
    credentials: true // Allows cookies to be sent
};

module.exports = cors(corsOptions);
