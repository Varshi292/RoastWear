import React, { useState, useEffect } from "react";
import { products } from "../Home/Product";
import { useDispatch } from "react-redux";
import { changeQuantity } from "../../stores/cart";

const CartItem = (props) => {
  const { productId, quantity } = props.data;
  const [detail, setDetail] = useState([]);
  const dispatch = useDispatch();

  useEffect(() => {
    const findDetail = products.filter(
      (product) => product.id === productId
    )[0];
    setDetail(findDetail);
  }, [productId]);

  const handleMinusQuantity = () => {
    dispatch(
      changeQuantity({
        productId: productId,
        quantity: quantity - 1,
      })
    );
  };

  const handlePlusQuantity = () => {
    dispatch(
      changeQuantity({
        productId: productId,
        quantity: quantity + 1,
      })
    );
  };

  return (
    <div className="flex justify-between items-center bg-slate-600 text-white p-2 border-b-2 border-slate-700 gap-5 rounded-md relative">
      <img src={detail.image} alt={detail.name} className="w-12" />

      <div className="flex-1 flex justify-between items-center gap-5">
        <h3 className="flex-1">{detail.name}</h3>
        {/* Price starts from 50% of the screen */}
        <p className="absolute left-1/2 transform -translate-x-1/2 text-center">
          {`$${detail.price * quantity}`}
        </p>
      </div>

      <div className="w-20 flex justify-between gap-2">
        <button
          className="bg-gray-200 rounded-full w-6 h-6 text-cyan-600"
          onClick={handleMinusQuantity}
        >
          -
        </button>
        <span>{quantity}</span>
        <button
          className="bg-gray-200 rounded-full w-6 h-6 text-cyan-600"
          onClick={handlePlusQuantity}
        >
          +
        </button>
      </div>
    </div>
  );
};

export default CartItem;
