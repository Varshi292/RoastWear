import React, { useState } from "react";
import { Link } from "react-router-dom";
import iconCart from "../../assets/images/iconCart.png";
import { useSelector, useDispatch } from "react-redux";
import { addToCart } from "../../stores/cart";

const ProductCart = (props) => {
  const carts = useSelector((store) => store.cart.items);
  const { id, name, price, designer, discount, image, slug } = props.data;
  const dispatch = useDispatch();
  const [isWishlisted, setIsWishlisted] = useState(false);
  const handleAddToCart = () => {
    dispatch(
      addToCart({
        productId: id,
        quantity: 1,
      })
    );
  };
  const toggleWishlist = () => {
    setIsWishlisted(!isWishlisted);
  };

  return (
    <div className="bg-white p-5 rounded-xl shadow-sm relative">
      <Link to={slug}>
        <img
          src={image}
          alt={name}
          className="product-image w-full h-64 object-cover transform transition duration-300 group-hover:scale-110"
        />
      </Link>
      <div className="absolute top-2 right-2">
        <button
          className="wishlist-btn bg-white m-5 p-2 rounded-full shadow-md hover:bg-red-100 transition duration-300"
          onClick={toggleWishlist}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill={isWishlisted ? "red" : "none"}
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className="w-6 h-6 text-red-500 hover:text-red-600"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.94l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78v0z"
            />
          </svg>
        </button>
      </div>

      <h3 className="text-lg sm:text-xl md:text-2xl font-semibold truncate">
        {name}
      </h3>
      <p className="text-gray-600 text-xs sm:text-sm">{designer}</p>

      <div className="flex justify-between items-center">
        <p className="price text-base sm:text-lg md:text-xl text-red-500">
          ${price.toFixed(2)}{" "}
          {discount && (
            <span className="text-gray-500 text-xs sm:text-sm">
              ({discount} off)
            </span>
          )}
        </p>
        <button
          className=" bg-blue-400 p-2 rounded-md text-sm hover:bg-gray-400 flex gap-2"
          onClick={handleAddToCart}
        >
          <img src={iconCart} alt="" className="w-5 " />
          Add To Cart
        </button>
      </div>
    </div>
  );
};

export default ProductCart;
