import React, { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import iconCart from "../../assets/images/iconCart.png";
import { useSelector, useDispatch } from "react-redux";
import { toggleStatusTab } from "../../stores/cart";
import "@fortawesome/fontawesome-free/css/all.css";
import CartTab from "../Carts/CartTab";
import { useSearch } from "../Context/SearchContext";

const Navbar = () => {
  const [menuOpen, setMenuOpen] = useState(false);
  const [totalQuantity, setTotalQuantity] = useState(0);
  const [search, setSearch] = useState("");
  const { searchTerm, setSearchTerm } = useSearch();


  const carts = useSelector((store) => store.cart.items);
  const dispatch = useDispatch();
  const navigate = useNavigate();

  useEffect(() => {
    let total = 0;
    carts.forEach((item) => (total += item.quantity));
    setTotalQuantity(total);
  }, [carts]);

  const toggleMenu = () => setMenuOpen(!menuOpen);

  const handleOpenTabCart = () => dispatch(toggleStatusTab());

  const handleSearch = (e) => {
    e.preventDefault();
    if (search.trim()) {
      navigate(`/shop?search=${encodeURIComponent(search.trim())}`);
      setSearch("");
    }
  };

  return (
    <nav className="bg-[#0b0c0f] text-gray-300 sticky top-0 z-50 shadow-md">
      <div className="container mx-auto flex items-center justify-between p-4">
        {/* Logo */}
        <div className="text-6xl font-extrabold tracking-wide relative">
          <Link to="/" className="text-[#25aae1] hover:text-[#1f8fcb] transition flex items-center group">
            R
            <span className="relative inline-block">
              oa
              <img
                src="/Assets/meme-glasses.png"
                alt="meme specs"
                className="absolute top-[-14px] left-[0px] w-120 pointer-events-none z-10 transform transition-all duration-500 ease-in-out group-hover:-translate-y-7 group-hover:rotate-[8deg]"
              />
            </span>
            stWear
          </Link>
        </div>

        {/* Navigation Links (Desktop) */}
        <div className="hidden md:flex items-center space-x-6 ml-6">
          <Link to="/" className="hover:text-[#ff2e63]">Home</Link>
          <Link to="/shop" className="hover:text-[#ff2e63]">Shop</Link>
          <Link to="/customize" className="hover:text-[#ff2e63]">Customize</Link>
          <Link to="/about" className="hover:text-[#ff2e63]">About</Link>
          <Link to="/login" className="hover:text-[#ff2e63]">Login</Link>
        </div>

        {/* Search */}
        <form onSubmit={handleSearch} className="hidden md:flex items-center space-x-2">
        <input
  type="text"
  placeholder="Search T-shirts..."
  className="p-2 rounded-md bg-[#1f2937] text-white placeholder:text-gray-400 focus:outline-none"
  value={searchTerm}
  onChange={(e) => setSearchTerm(e.target.value)}
/>

          <button type="submit">
            <i className="fas fa-search text-[#25aae1]" />
          </button>
        </form>

        {/* Wishlist + Cart (Desktop) */}
        <div className="hidden md:flex items-center space-x-4">
          <Link to="/wishlist" className="hover:text-[#ff2e63]">
            <i className="fas fa-heart mr-1"></i> Wishlist
          </Link>

          <div
            className="w-10 h-10 bg-[#f3efe9] rounded-full flex justify-center items-center relative cursor-pointer"
            onClick={handleOpenTabCart}
          >
            <img src={iconCart} alt="Cart" className="w-6" />
            <span className="absolute top-2/3 right-1/2 bg-[#ff2e63] text-white text-sm w-5 h-5 rounded-full flex justify-center items-center">
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
          menuOpen ? "fixed inset-0 bg-[#121417] bg-opacity-95 z-50" : "hidden"
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
        <Link to="/" onClick={toggleMenu} className="hover:text-[#25aae1]">Home</Link>
        <Link to="/shop" onClick={toggleMenu} className="hover:text-[#25aae1]">Shop</Link>
        <Link to="/customize" onClick={toggleMenu} className="hover:text-[#25aae1]">Customize</Link>
        <Link to="/about" onClick={toggleMenu} className="hover:text-[#25aae1]">About</Link>
        <Link to="/login" onClick={toggleMenu} className="hover:text-[#25aae1]">Login</Link>
        <Link to="/wishlist" onClick={toggleMenu} className="hover:text-[#ff2e63]">
          <i className="far fa-heart"></i> Wishlist
        </Link>
        <Link to="/cart" onClick={toggleMenu} className="hover:text-[#ff2e63]">
          <i className="fas fa-shopping-bag"></i> Cart
        </Link>
      </div>

      {/* Cart Tab */}
      <CartTab />
    </nav>
  );
};

export default Navbar;
