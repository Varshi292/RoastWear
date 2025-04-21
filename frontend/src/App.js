// src/App.js
import React from "react";
import { Routes, Route, useLocation } from "react-router-dom";
import "./index.css";

import Navbar from "./components/Navbar/Navbar";
import Home from "./components/Home/Home";
import Shop from "./components/Shop/Shop";
import Customize from "./components/Customize/Customize";
import About from "./components/About/About";
import Contact from "./components/Contact/Contact";
import Login from "./components/Login/Login";
import Register from "./components/Login/Register";
import Detail from "./components/Home/detail";
import WishlistPage from "./Pages/WishlistPage";
import WelcomeBanner from "./components/Banners/WelcomeBanner";
import { useUser } from "./components/Context/UserContext";

function App() {
  const { userName } = useUser();
  const location = useLocation(); 

  const isHomeRoute = location.pathname === "/" || location.pathname === "/home";

  return (
    <>
      <Navbar />

   
      {userName && isHomeRoute && <WelcomeBanner />}

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/home" element={<Home />} />
        <Route path="/shop" element={<Shop />} />
        <Route path="/customize" element={<Customize />} />
        <Route path="/about" element={<About />} />
        <Route path="/contact" element={<Contact />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/product/:slug" element={<Detail />} />
        <Route path="/wishlist" element={<WishlistPage />} />
      </Routes>
    </>
  );
}

export default App;
