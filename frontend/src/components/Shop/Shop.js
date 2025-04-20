// src/components/Shop/Shop.js
import React from "react";
import { products } from "../Home/Product";
import ProductCart from "../Carts/productCart";
import { useSearch } from "../Context/SearchContext";

const Shop = () => {
  const { searchTerm } = useSearch(); // get from context

  const filteredProducts = products.filter((product) =>
    product.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="bg-[#0b0c0f] min-h-screen text-gray-100 px-6 py-12">
      <h1 className="text-3xl md:text-4xl font-bold text-center mb-12 text-[#25aae1] drop-shadow-[0_0_8px_#25aae1]">
        🛍️ Shop All T-Shirts
      </h1>

      <div className="grid gap-10 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        {filteredProducts.length > 0 ? (
          filteredProducts.map((product) => (
            <ProductCart key={product.id} data={product} />
          ))
        ) : (
          <p className="text-center col-span-full text-gray-400">
            No products found matching "<strong>{searchTerm}</strong>"
          </p>
        )}
      </div>
    </div>
  );
};

export default Shop;
