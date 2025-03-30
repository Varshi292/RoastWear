// src/components/Home/detail.js
import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { products } from "./Product";
import { useDispatch } from "react-redux";
import { addToCart } from "../../stores/cart";

const Detail = () => {
  const { slug } = useParams();
  const navigate = useNavigate();
  const [product, setProduct] = useState(null);
  const [quantity, setQuantity] = useState(1);
  const dispatch = useDispatch();

  useEffect(() => {
    const match = products.find((item) => item.slug === slug);
    if (match) {
      setProduct(match);
    } else {
      navigate("/"); // Redirect if slug is invalid
    }
  }, [slug, navigate]);

  const handleMinusQuantity = () => {
    setQuantity((prev) => (prev > 1 ? prev - 1 : 1));
  };

  const handlePlusQuantity = () => {
    setQuantity((prev) => prev + 1);
  };

  const handleAddToCart = () => {
    if (product) {
      dispatch(
        addToCart({
          productId: product.id,
          quantity,
        })
      );
    }
  };

  if (!product) return null;

  return (
    <div className="flex justify-center items-center min-h-screen px-4 bg-gray-100">
      <div className="bg-white p-6 rounded-lg shadow-lg max-w-6xl w-full">
        <h2 className="text-3xl font-semibold text-center mb-12">
          PRODUCT DETAIL
        </h2>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-10">
          {/* Image */}
          <div className="flex justify-center">
            <img
              src={product.image}
              alt={product.name}
              className="w-72 h-72 object-cover rounded-lg shadow-md"
            />
          </div>

          {/* Details */}
          <div className="flex flex-col gap-5 justify-center">
            <h1 className="text-2xl font-bold">{product.name}</h1>
            <p className="text-xl text-red-500 font-semibold">
              ${product.price.toFixed(2)}
            </p>
            <p className="text-sm text-gray-700">
              {product.description || "No description provided."}
            </p>

            {/* Quantity + Add to Cart */}
            <div className="flex gap-6 items-center mt-6">
              <div className="flex items-center border rounded overflow-hidden">
                <button
                  onClick={handleMinusQuantity}
                  className="bg-gray-100 px-4 py-2 text-lg font-bold"
                >
                  −
                </button>
                <span className="px-5 py-2 bg-gray-200">{quantity}</span>
                <button
                  onClick={handlePlusQuantity}
                  className="bg-gray-100 px-4 py-2 text-lg font-bold"
                >
                  +
                </button>
              </div>

              <button
                onClick={handleAddToCart}
                className="bg-slate-900 text-white px-6 py-2 rounded shadow-md hover:bg-slate-700 transition duration-300"
              >
                Add to Cart
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Detail;
