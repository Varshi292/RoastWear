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
    <button onClick={handleToggle} aria-label="Toggle wishlist">
      {isInWishlist ? (
        <FaHeart className="text-red-500 text-xl" />
      ) : (
        <FaRegHeart className="text-gray-400 hover:text-red-500 text-xl transition-all" />
      )}
    </button>
  );
};

export default WishlistButton;
