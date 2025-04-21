import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const Register = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleRegister = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch("http://localhost:7777/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, email, password }),
      });

      const result = await response.json();
      setMessage(result.message);

      if (result.success) {
        // ✅ Registration successful, redirect to login
        setTimeout(() => navigate("/login"), 1500);
      }
    } catch (error) {
      console.error("Registration error:", error);
      setMessage("Registration failed. Try again.");
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <form
        onSubmit={handleRegister}
        className="bg-white p-8 rounded-md shadow-md w-full max-w-sm"
      >
        <h2 className="text-2xl font-bold mb-4 text-center">Register</h2>

        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className="block w-full p-2 border border-gray-300 rounded mt-2"
          required
        />

        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="block w-full p-2 border border-gray-300 rounded mt-4"
          required
        />

        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="block w-full p-2 border border-gray-300 rounded mt-4"
          required
        />

        <button
          type="submit"
          className="w-full bg-green-500 hover:bg-green-600 text-white p-2 rounded mt-4"
        >
          Register
        </button>

        {message && (
          <p className="text-sm mt-3 text-center text-green-600">{message}</p>
        )}

        <p className="text-sm mt-4 text-center">
          Already have an account?{" "}
          <span
            onClick={() => navigate("/login")}
            className="text-blue-500 underline cursor-pointer"
          >
            Login
          </span>
        </p>
      </form>
    </div>
  );
};

export default Register;
