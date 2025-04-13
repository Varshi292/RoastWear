import React from "react";
import { useDispatch, useSelector } from "react-redux";
import { toggleWishlist } from "../../stores/wishlist";
import { FaHeart, FaRegHeart } from "react-icons/fa";

const WishlistButton = ({ productId }) => {
  const dispatch = useDispatch();
  const wishlist = useSelector((state) => state.wishlist.items);
  const isInWishlist = wishlist.includes(productId);

  const handleToggle = () => {
    dispatch(toggleWishlist(productId));
  };

  return (
    <button
      onClick={handleToggle}
      aria-label="Toggle wishlist"
      className="transition-all hover:scale-110"
    >
      {isInWishlist ? (
        <FaHeart className="text-pink-500 text-2xl drop-shadow-[0_0_5px_#ff2e63]" />
      ) : (
        <FaRegHeart className="text-gray-500 hover:text-pink-400 text-2xl transition-colors duration-200" />
      )}
    </button>
  );
};

export default WishlistButton;
