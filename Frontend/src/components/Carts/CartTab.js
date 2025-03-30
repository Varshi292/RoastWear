import React from "react";
import { useSelector, useDispatch } from "react-redux";
import CartItem from "./CartItem";
import { toggleStatusTab } from "../../stores/cart";

const CartTab = () => {
  const carts = useSelector((store) => store.cart.items);
  const statusTab = useSelector((store) => store.cart.statusTab);
  const dispatch = useDispatch();

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

  if (!statusTab) return null;

  return (
    <div className="fixed inset-0 z-50 bg-zinc-50 text-zinc-900 grid grid-rows-[80px_1fr_80px]">
      {/* Header */}
      <header
        className="flex items-center justify-between px-6 bg-blue-600
 text-white shadow-xl"
      >
        <h2 className="text-2xl font-semibold tracking-wide">🛒 Your Cart</h2>
        <button
          onClick={handleCloseTabCart}
          className="text-rose-500 hover:text-rose-700 text-2xl font-bold transition-all"
          aria-label="Close cart"
        >
          ✕
        </button>
      </header>

      {/* Cart Content */}
      <main className="overflow-y-auto px-6 py-4 bg-zinc-50">
        {carts.length > 0 ? (
          carts.map((item, idx) => <CartItem key={idx} data={item} />)
        ) : (
          <p className="text-center text-zinc-500 mt-10 text-lg">
            Your cart is currently empty 🛍️
          </p>
        )}
      </main>

      {/* Footer */}
      <footer className="bg-white border-t border-zinc-200 shadow-inner px-6 flex items-center justify-between">
        <span className="text-xl font-bold text-zinc-800">
          Total: ${getTotalPrice().toFixed(2)}
        </span>
        <button
          className="bg-blue-600
 text-white px-6 py-2 rounded-lg font-medium hover:bg-green-600 transition"
        >
          Proceed to Checkout
        </button>
      </footer>
    </div>
  );
};

export default CartTab;
