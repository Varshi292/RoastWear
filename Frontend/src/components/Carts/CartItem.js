import React, { useState, useEffect } from "react";
import { products } from "../Home/Product";
import { useDispatch } from "react-redux";
import { changeQuantity, removeFromCart } from "../../stores/cart";
import { FaMinus, FaPlus, FaTrash } from "react-icons/fa";

const CartItem = ({ data }) => {
  const { productId, quantity } = data;
  const [detail, setDetail] = useState({});
  const dispatch = useDispatch();

  useEffect(() => {
    const findDetail = products.find((product) => product.id === productId);
    setDetail(findDetail || {});
  }, [productId]);

  const handleMinusQuantity = () => {
    if (quantity > 1) {
      dispatch(changeQuantity({ productId, quantity: quantity - 1 }));
    }
  };

  const handlePlusQuantity = () => {
    dispatch(changeQuantity({ productId, quantity: quantity + 1 }));
  };
  const handleRemove = () => {
    dispatch(removeFromCart(productId));
  };

  if (!detail) return null;

  return (
    <div className="flex items-center bg-white text-zinc-900 p-4 rounded-xl shadow-lg mb-4 transition-all duration-300 hover:shadow-xl border border-zinc-200">
      {/* Product Image */}
      <img
        src={detail.image}
        alt={detail.name}
        className="w-20 h-20 object-cover rounded-md"
      />

      {/* Info Section */}
      <div className="ml-4 flex-1">
        <h3 className="text-lg font-semibold">{detail.name}</h3>
        <p className="text-sm text-zinc-500">{detail.designer}</p>
        <p className="text-lg font-bold text-indigo-700 mt-1">
          ${(detail.price * quantity).toFixed(2)}
        </p>
      </div>

      {/* Quantity + Remove Controls */}
      <div className="flex items-center gap-3">
        <div className="flex items-center gap-2 bg-zinc-100 rounded-full px-3 py-1">
          <button
            onClick={handleMinusQuantity}
            aria-label="minus"
            className="text-rose-500 hover:text-rose-700 transition"
          >
            <FaMinus />
          </button>
          <span className="px-2 font-medium">{quantity}</span>
          <button
            onClick={handlePlusQuantity}
            aria-label="plus"
            className="text-green-600 hover:text-green-800 transition"
          >
            <FaPlus />
          </button>
        </div>
        {/* Trash icon on the same row */}
        <button
          onClick={handleRemove}
          className="text-zinc-400 hover:text-rose-600 transition"
          aria-label="Remove item"
        >
          <FaTrash />
        </button>
      </div>
    </div>
  );
};

export default CartItem;
