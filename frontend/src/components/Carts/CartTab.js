import React from "react";
import { useSelector, useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import CartItem from "./CartItem"; // adjust path if needed
import { toggleStatusTab, clearCart } from "../../stores/cart";


const CartTab = () => {
  const carts = useSelector((store) => store.cart.items);
  const statusTab = useSelector((store) => store.cart.statusTab);
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const handleCloseTabCart = () => {
    dispatch(toggleStatusTab());
  };

  const getTotalPrice = () => {
    const { products } = require("../Home/Product");
    return carts.reduce((total, item) => {
      const product = products.find((p) => p.id === item.productId);
      return total + (product ? product.price * item.quantity : 0);
    }, 0);
  };

  const purchaseItems = async () => {
    const username = localStorage.getItem("userName");
  
    if (!username) {
      alert("You must be logged in to complete the purchase.");
      return;
    }
  
    try {
      const response = await fetch("http://localhost:7777/checkout", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username }),
      });
  
      const result = await response.json();
      if (result.success) {
        console.log("✅ Checkout successful:", result.message);
        dispatch(clearCart()); // ← clear frontend cart
      } else {
        console.error("❌ Checkout failed:", result.error);
      }
    } catch (err) {
      console.error("❌ Error during checkout:", err);
    }
  };
  
  

  if (!statusTab) return null;

  return (
    <div className="fixed inset-0 z-50 bg-[#0b0c0f] text-gray-300 grid grid-rows-[80px_1fr_80px]">
      {/* Header */}
      <header className="flex items-center justify-between px-6 bg-[#25aae1] text-white shadow-xl">
        <h2 className="text-2xl font-semibold tracking-wide">🛒 Your Cart</h2>
        <button
          onClick={handleCloseTabCart}
          className="text-white hover:text-rose-400 text-2xl font-bold transition-all"
          aria-label="Close cart"
        >
          ✕
        </button>
      </header>

      {/* Cart Content */}
      <main className="overflow-y-auto px-6 py-4 bg-[#121417]">
        {carts.length > 0 ? (
          carts.map((item, idx) => <CartItem key={idx} data={item} />)
        ) : (
          <p className="text-center text-gray-500 mt-10 text-lg">
            Your cart is currently empty 🛍️
          </p>
        )}
      </main>

      {/* Footer */}
      <footer className="bg-[#1f1f1f] border-t border-gray-700 shadow-inner px-6 flex items-center justify-between">
        <span className="text-xl font-bold text-gray-100">
          Total: ${getTotalPrice().toFixed(2)}
        </span>
        <button
          className="bg-[#25aae1] text-white px-6 py-2 rounded-lg font-medium hover:bg-[#1f8fcb] transition"
          onClick={purchaseItems}
        >
          Proceed to Checkout
        </button>
      </footer>
    </div>
  );
};

export default CartTab;
