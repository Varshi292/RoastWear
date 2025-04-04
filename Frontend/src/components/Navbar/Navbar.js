import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import iconCart from "../../assets/images/iconCart.png";
import { useSelector, useDispatch } from "react-redux";
import { toggleStatusTab } from "../../stores/cart";
import "@fortawesome/fontawesome-free/css/all.css";
import CartTab from "../Carts/CartTab";

const Navbar = () => {
  const [menuOpen, setMenuOpen] = useState(false);
  const [totalQuantity, setTotalQuantity] = useState(0);

  const carts = useSelector((store) => store.cart.items);
  const dispatch = useDispatch();

  useEffect(() => {
    let total = 0;
    carts.forEach((item) => (total += item.quantity));
    setTotalQuantity(total);
  }, [carts]);

  const toggleMenu = () => setMenuOpen(!menuOpen);

  const handleOpenTabCart = () => dispatch(toggleStatusTab());

  return (
    <nav className="bg-gray-800 text-white sticky top-0 z-50">
      <div className="container mx-auto flex items-center justify-between p-4">
        {/* Logo */}
        <div className="text-2xl font-bold">
          <Link to="/" className="text-white">
            T-Shirt Customizer
          </Link>
        </div>

        {/* Navigation Links (Desktop) */}
        <div className="hidden md:flex items-center space-x-6 ml-6">
          <Link to="/" className="hover:text-gray-400">Home</Link>
          <Link to="/shop" className="hover:text-gray-400">Shop</Link>
          <Link to="/customize" className="hover:text-gray-400">Customize</Link>
          <Link to="/about" className="hover:text-gray-400">About</Link>
          <Link to="/login" className="hover:text-gray-400">Login</Link>
        </div>

        {/* Search */}
        <div className="hidden md:flex items-center space-x-2">
          <input
            type="text"
            placeholder="Search T-shirts..."
            className="p-2 rounded-md bg-gray-700 text-white"
          />
          <i className="fas fa-search"></i>
        </div>

        {/* Wishlist + Cart (Desktop) */}
        <div className="hidden md:flex items-center space-x-4">
        <Link to="/wishlist" className="hover:text-red-500">
  <i className="fas fa-heart mr-1"></i> Wishlist
</Link>

          <div
            className="w-10 h-10 bg-gray-100 rounded-full flex justify-center items-center relative"
            onClick={handleOpenTabCart}
          >
            <img src={iconCart} alt="Cart" className="w-6" />
            <span className="absolute top-2/3 right-1/2 bg-red-500 text-white text-sm w-5 h-5 rounded-full flex justify-center items-center">
              {totalQuantity}
            </span>
          </div>
        </div>

        {/* Hamburger Menu (Mobile) */}
        <button
          className="md:hidden text-2xl focus:outline-none"
          aria-label="open menu"
          onClick={toggleMenu}
        >
          <i className="fas fa-bars" />
        </button>
      </div>

      {/* Mobile Menu */}
      <div
        className={`${
          menuOpen ? "fixed inset-0 bg-gray-900 bg-opacity-95 z-50" : "hidden"
        } flex flex-col items-center justify-center space-y-8 text-2xl`}
      >
        {/* Close button */}
        <button
          className="absolute top-4 right-4 text-3xl text-white"
          aria-label="close menu"
          onClick={toggleMenu}
        >
          <i className="fas fa-times" />
        </button>

        {/* Mobile Links */}
        <Link to="/" onClick={toggleMenu} className="hover:text-gray-400">Home</Link>
        <Link to="/shop" onClick={toggleMenu} className="hover:text-gray-400">Shop</Link>
        <Link to="/customize" onClick={toggleMenu} className="hover:text-gray-400">Customize</Link>
        <Link to="/about" onClick={toggleMenu} className="hover:text-gray-400">About</Link>
        <Link to="/login" onClick={toggleMenu} className="hover:text-gray-400">Login</Link>
        <Link to="/wishlist" onClick={toggleMenu} className="hover:text-gray-400">
          <i className="far fa-heart"></i> Wishlist
        </Link>
        <Link to="/cart" onClick={toggleMenu} className="hover:text-gray-400">
          <i className="fas fa-shopping-bag"></i> Cart
        </Link>
      </div>

      {/* Cart Tab */}
      <CartTab />
    </nav>
  );
};

export default Navbar;
