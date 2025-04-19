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
  return (
    <div
      className={`fixed top-0 right-0 bg-gray-700 shadow-2xl w-full h-full grid grid-rows-[60px_1fr_60px] 
    transform transition-transform duration-500
    ${statusTab === false ? "translate-x-full" : ""}
    `}
    >
      <h2 className="p-5 text-white text-2xl">Shopping Cart</h2>
      <div className="p-5">
        {carts.map((item, key) => (
          <CartItem key={key} data={item} />
        ))}
      </div>
      <div className="grid grid-cols-2">
        <button
          className="bg-red-500 text-black py-2 px-6 rounded-md hover:bg-red-800 hover:text-white transition-all duration-300"
          onClick={handleCloseTabCart}
        >
          CLOSE
        </button>
        <button className="bg-blue-400 text-black hover:bg-green-800 hover:text-white">
          CHECKOUT
        </button>
      </div>
    </div>
  );
};

export default CartTab;
