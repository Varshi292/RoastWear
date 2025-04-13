// src/pages/WishlistPage.js
import React from "react";
import { useSelector } from "react-redux";
import { products } from "../components/Home/Product";
import ProductCart from "../components/Carts/productCart";

const WishlistPage = () => {
  const wishlistItems = useSelector((state) => state.wishlist.items);

  const wishlistProducts = products.filter((product) =>
    wishlistItems.includes(product.id)
  );

  return (
    <div className="min-h-screen bg-[#0b0c0f] text-gray-300 p-4 md:p-8">
      <h1 className="text-3xl font-bold mb-6 text-center text-pink-500 drop-shadow-[0_0_5px_#ff2e63]">
        Your Wishlist 
      </h1>

      {wishlistProducts.length === 0 ? (
        <p className="text-center text-gray-500 text-lg">
          Your wishlist is empty.
        </p>
      ) : (
        <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {wishlistProducts.map((product) => (
            <ProductCart key={product.id} data={product} />
          ))}
        </div>
      )}
    </div>
  );
};

export default WishlistPage;
