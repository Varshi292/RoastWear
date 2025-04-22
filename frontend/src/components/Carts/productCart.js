import React from "react";
import { Link } from "react-router-dom";
import iconCart from "../../assets/images/iconCart.png";
import { useSelector, useDispatch } from "react-redux";
import { addToCart } from "../../stores/cart";
import { toggleWishlist } from "../../stores/wishlist";
import { FaHeart, FaRegHeart } from "react-icons/fa";

const ProductCart = ({ data }) => {
  const dispatch = useDispatch();

  const wishlistItems = useSelector((state) => state.wishlist?.items || []);

  const { id, name, price, designer, discount, image, slug } = data;

  const isWishlisted = Array.isArray(wishlistItems)
    ? wishlistItems.includes(id)
    : false;

    const handleAddToCart = async () => {
      const quantity = 1;
      const username = localStorage.getItem("userName");
      const sessionid = "placeholder"; // or use actual session ID logic if applicable
    
      // Update UI immediately
      dispatch(
        addToCart({
          productId: id,
          quantity,
          price,
          username,
          sessionid,
        })
      );
    };
    
    


  const handleToggleWishlist = () => {
    dispatch(toggleWishlist(id));
  };

  return (
    <div
      data-testid="product-card"
      className="bg-[#1e1e1e] p-5 rounded-xl shadow-[0_0_10px_rgba(255,255,255,0.05)] relative text-gray-300 border border-[#2c2c2c]"
    >
      <Link to={`/product/${slug}`}>
        <img
          src={image}
          alt={name}
          className="product-image w-full h-64 object-cover transform transition duration-300 rounded-lg hover:scale-105"
        />
      </Link>

      {/* Wishlist Button */}
      <div className="absolute top-2 right-2">
        <button
          className="bg-[#1e1e1e] m-2 p-2 rounded-full shadow-md hover:bg-[#2a2a2a] transition duration-300"
          onClick={handleToggleWishlist}
          aria-label="Toggle Wishlist"
          data-testid={`wishlist-toggle-${id}`}
        >
          {isWishlisted ? (
            <FaHeart className="text-red-500 w-6 h-6" />
          ) : (
            <FaRegHeart className="text-gray-500 hover:text-red-500 w-6 h-6" />
          )}
        </button>
      </div>

      <h3 className="text-lg sm:text-xl md:text-2xl font-semibold truncate text-gray-300">
        {name}
      </h3>

      <p className="text-gray-500 text-xs sm:text-sm">{designer}</p>

      <div className="flex justify-between items-center mt-2">
        <p className="price text-base sm:text-lg md:text-xl text-[#ff2e63]">
          ${price.toFixed(2)}{" "}
          {discount && (
            <span className="text-gray-500 text-xs sm:text-sm">
              ({discount} off)
            </span>
          )}
        </p>
        <button
          className="bg-[#25aae1] text-white p-2 rounded-md text-sm hover:bg-[#1f8fcb] flex items-center gap-2 shadow hover:shadow-md"
          onClick={handleAddToCart}
          aria-label="Add to Cart"
          data-testid={`add-to-cart-${id}`}
        >
          <img src={iconCart} alt="Cart Icon" className="w-5" />
        </button>
      </div>

      <Link to={`/product/${slug}`}>
        <button
          className="mt-4 w-full bg-[#25aae1] hover:bg-[#1f8fcb] text-black font-semibold py-2 rounded"
          data-testid="view-button"
        >
          View Product
        </button>
      </Link>
    </div>
  );
};

export default ProductCart;
