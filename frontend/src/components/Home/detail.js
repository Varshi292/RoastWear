import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { products } from "./Product";
import { useDispatch } from "react-redux";
import { addToCart } from "../../stores/cart";

const Detail = () => {
  const { slug } = useParams();
  const [detail, setDetail] = useState([]);
  const [quantity, setQuantity] = useState(1);
  const dispatch = useDispatch();

  useEffect(() => {
    const findDetail = products.filter((product) => product.slug === slug);
    if (findDetail.length > 0) {
      setDetail(findDetail[0]);
    } else {
      window.location.href = "/";
    }
  }, [slug]);

  const handleMinusQuantity = () => {
    setQuantity(quantity - 1 < 1 ? 1 : quantity - 1);
  };

  const handlePlusQuantity = () => {
    setQuantity(quantity + 1);
  };

  const handleAddToCart = () => {
    dispatch(
      addToCart({
        productId: detail.id,
        quantity: quantity,
      })
    );
  };

  return (
    <div className="flex justify-center items-center min-h-screen px-2">
      <div className="bg-white p-6 rounded-lg shadow-lg max-w-6xl w-full h-auto min-h-[600px]">
        <h2 className="text-3xl font-semibold text-center mb-20">
          PRODUCT DETAIL
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-2">
          {/* Product Image */}
          <div className="flex justify-center">
            <img
              src={detail.image}
              alt={detail.name}
              className="w-72 h-72 object-cover rounded-lg shadow-md"
            />
          </div>

          {/* Product Details */}
          <div className="flex flex-col gap-4">
            <h1 className="text-1xl font-bold uppercase">{detail.name}</h1>
            <p className="text-xl font-semibold">${detail.price}</p>
            <p className="text-sm text-gray-700">{detail.description}</p>

            {/* Quantity Selector */}
            <div className="flex gap-20 items-center mt-6">
              <div className="flex items-center border rounded-lg overflow-hidden">
                <button
                  className="bg-gray-100 px-3 py-2 font-bold text-lg"
                  onClick={handleMinusQuantity}
                >
                  -
                </button>
                <span className="px-4 py-2 bg-gray-200">{quantity}</span>
                <button
                  className="bg-gray-100 px-3 py-2 font-bold text-lg"
                  onClick={handlePlusQuantity}
                >
                  +
                </button>
              </div>

              {/* Add to Cart Button */}

              <button
                className="bg-slate-900 text-white px-6 py-2 rounded-lg shadow-md hover:bg-gray-800 transition duration-300"
                onClick={handleAddToCart}
              >
                Add To Cart
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Detail;
