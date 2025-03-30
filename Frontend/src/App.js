// src/App.js
import React from "react";
import "./index.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Navbar from "./components/Navbar/Navbar";
import Home from "./components/Home/Home";
import Shop from "./components/Shop/Shop";
import Customize from "./components/Customize/Customize";
import About from "./components/About";
import Contact from "./components/Contact";
import Login from "./components/Login/Login";
import Register from "./components/Login/Register";
import Detail from "./components/Home/detail";
import WishlistPage from "./Pages/WishlistPage";




// import WishlistPage from "./pages/WishlistPage";

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/shop" element={<Shop />} />
        <Route path="/customize" element={<Customize />} />
        <Route path="/about" element={<About />} />
        <Route path="/contact" element={<Contact />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/home" element={<Home />} />
        <Route path="/product/:slug" element={<Detail />} />
        <Route path="/wishlist" element={<WishlistPage />} /> 
      </Routes>
    </Router>
  );
}


export default App;
