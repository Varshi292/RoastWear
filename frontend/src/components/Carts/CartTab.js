import React from "react";
import { useSelector, useDispatch } from "react-redux";
import CartItem from "./CartItem";
import {
  toggleStatusTab,
  clearCart,
} from "../../stores/cart";

const CartTab = () => {
  const carts = useSelector((store) => store.cart.items);
  const statusTab = useSelector((store) => store.cart.statusTab);
  const dispatch = useDispatch();

  const handleCloseTabCart = () => {
    console.log("Closing tab");
    dispatch(toggleStatusTab());
  };

  const handleCheckout = async () => {
    console.log("comfirming checkout");
    const username = localStorage.getItem("username");
    const session_id = localStorage.getItem("session_id");

    if (!username || !session_id) {
      alert("⚠️ You're not logged in or session expired.");
      return;
    }

    try {
      const response = await fetch("http://localhost:7777/checkout", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          username,
          session_id,
          items: carts,
        }),
      });

      const result = await response.json();

      if (response.ok) {
        alert("✅ Checkout successful!");

        // ✅ Clear cart from Redux and localStorage
        dispatch(clearCart());
        localStorage.setItem("carts", JSON.stringify([]));

        handleCloseTabCart(); // Close the cart tab
      } else {
        alert(`❌ Checkout failed: ${result.error || "Unknown error"}`);
      }
    } catch (err) {
      console.error(err);
      alert("❌ Error during checkout. Please try again.");
    }
  };

  return (
    <div
      className={`fixed top-0 right-0 bg-gray-700 shadow-2xl w-full h-full grid grid-rows-[60px_1fr_60px] 
      transform transition-transform duration-500
      ${statusTab === false ? "translate-x-full" : ""}
    `}
    >
      <h2 className="p-5 text-white text-2xl">Shopping Cart</h2>
      <div className="p-5 overflow-y-auto">
        {carts.map((item, key) => (
          <CartItem key={key} data={item} />
        ))}
        {carts.length === 0 && (
          <p className="text-white mt-4 text-center">Your cart is empty.</p>
        )}
      </div>
      <div className="grid grid-cols-2">
        <button
          className="bg-red-500 text-black py-2 px-6 rounded-md hover:bg-red-800 hover:text-white transition-all duration-300"
          onClick={handleCloseTabCart}
        >
          CLOSE
        </button>
        <button
          className="bg-blue-400 text-black hover:bg-green-800 hover:text-white transition-all duration-300"
          onClick={handleCheckout}
          disabled={carts.length === 0}
        >
          CHECKOUT
        </button>
      </div>
    </div>
  );
};

export default CartTab;
