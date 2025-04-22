// src/components/Login/Login.js
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useUser } from "../Context/UserContext";
import { products } from "../Home/Product";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const [isSessionChecked, checkSession] = useState(false);
  const navigate = useNavigate();
  const { setUserName } = useUser();

  useEffect(() => {
    const verifySession = async () => {
      const response = await fetch("http://localhost:7777/session/verify", {
        method: "GET",
        credentials: "include",
      });
      const result = await response.json();
      if (result.success) {
        localStorage.setItem("userName", username);
        setUserName(username);
        navigate("/");
      } else {
        checkSession(true);
      }
    };
    verifySession();
  }, [navigate]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch("http://localhost:7777/login", {
        method: "POST",
        credentials: "include",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      const result = await response.json();
      setMessage(result.message);

      if (result.success) {
        localStorage.setItem("userName", username);
        setUserName(username);

        const localCart = JSON.parse(localStorage.getItem("carts")) || [];

        const backendResponse = await fetch(`http://localhost:7777/cart/items?username=${username}`, {
          credentials: "include",
        });
        const backendCart = await backendResponse.json();

        const mergedCartMap = new Map();

        backendCart.forEach(item => {
          const product = products.find(p => p.id === item.ProductID);
          const price = product ? product.price : 0;

          mergedCartMap.set(item.ProductID, {
            productId: item.ProductID,
            quantity: item.Quantity,
            price,
          });
        });

        localCart.forEach(item => {
          const product = products.find(p => p.id === item.productId);
          const price = product ? product.price : 0;

          if (mergedCartMap.has(item.productId)) {
            const existing = mergedCartMap.get(item.productId);
            mergedCartMap.set(item.productId, {
              ...existing,
              quantity: existing.quantity + item.quantity,
            });
          } else {
            mergedCartMap.set(item.productId, {
              productId: item.productId,
              quantity: item.quantity,
              price: price,
            });
          }
        });

        for (const item of mergedCartMap.values()) {
          await fetch("http://localhost:7777/cart/modify", {
            method: "POST",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              username,
              sessionid: "placeholder",
              productid: item.productId,
              quantity: item.quantity,
              unitPrice: item.price * item.quantity, // unitPrice only
            }),
          });
        }

        navigate("/");
      }
    } catch (error) {
      setMessage("Login failed. Please try again.");
      console.error(error);
    }
  };

  if (!isSessionChecked) return null;

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <form className="bg-white p-8 rounded-md shadow-md" onSubmit={handleSubmit}>
        <h2 className="text-2xl font-bold mb-4">Login</h2>

        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className="block w-full p-2 border border-gray-300 rounded mt-2"
        />

        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="block w-full p-2 border border-gray-300 rounded mt-2"
        />

        <button
          type="submit"
          className="w-full bg-blue-500 hover:bg-blue-600 text-white p-2 rounded mt-4"
        >
          Login
        </button>

        {message && (
          <p
            className={`mt-2 text-sm ${message.includes("success") ? "text-green-600" : "text-red-600"}`}
          >
            {message}
          </p>
        )}

        <p className="text-sm mt-4">
          New user?{" "}
          <span
            onClick={() => navigate("/register")}
            className="text-blue-500 underline cursor-pointer"
          >
            Register here
          </span>
        </p>
      </form>
    </div>
  );
};

export default Login;
