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
      navigate("/");
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
    <div className="flex justify-center items-center min-h-screen px-4 bg-[#0b0c0f] text-gray-200 pb-4">
      <div className="bg-[#1e1e1e] p-10 rounded-xl shadow-xl max-w-6xl w-full border border-[#2c2c2c] min-h-[500px]">
        

        <div className="grid grid-cols-1 md:grid-cols-2 gap-10">
          <div className="flex justify-center">
            <img
              src={product.image}
              alt={product.name}
              className="w-72 h-72 object-cover rounded-lg shadow-md"
            />
          </div>

          <div className="flex flex-col gap-5 justify-center">
            <h1 className="text-3xl font-bold text-white">{product.name}</h1>
            <p className="text-xl text-[#ff2e63] font-semibold">
              ${product.price.toFixed(2)}
            </p>
            <p className="text-sm text-gray-400">
              {product.description || "No description provided."}
            </p>

            <div className="flex gap-6 items-center mt-6">
              <div className="flex items-center border border-[#2c2c2c] rounded overflow-hidden">
                <button
                  onClick={handleMinusQuantity}
                  className="bg-[#2c2c2c] px-4 py-2 text-lg font-bold text-white hover:bg-[#3b3b3b]"
                >
                  −
                </button>
                <span className="px-5 py-2 bg-[#3b3b3b] text-white">{quantity}</span>
                <button
                  onClick={handlePlusQuantity}
                  className="bg-[#2c2c2c] px-4 py-2 text-lg font-bold text-white hover:bg-[#3b3b3b]"
                >
                  +
                </button>
              </div>

              <button
                onClick={handleAddToCart}
                className="bg-[#22a7e0] text-white px-6 py-2 rounded hover:bg-[#1a9fd5] shadow-md transition duration-300"
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
