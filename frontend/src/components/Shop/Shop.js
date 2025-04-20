import React from "react";
import { products } from "../Home/Product";
import ProductCart from "../Carts/productCart"; 

const Shop = () => {
  return (
    <div className="bg-[#0b0c0f] min-h-screen text-gray-100 px-6 py-12">
      <h1 className="text-3xl md:text-4xl font-bold text-center mb-12 text-[#25aae1] drop-shadow-[0_0_8px_#25aae1]">
        🛍️ Shop All T-Shirts
      </h1>

      <div className="grid gap-10 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        {products.map((product) => (
          <ProductCart key={product.id} data={product} />
        ))}
      </div>
    </div>
  );
};

export default Shop;
