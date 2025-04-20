import React from "react";
import { products } from "../Home/Product";
import { Link } from "react-router-dom";
import { useSearch } from "../Context/SearchContext";

const Shop = () => {
  const { searchTerm } = useSearch();

  const filteredProducts = products.filter((product) =>
    product.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    product.designer.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="bg-[#0b0c0f] min-h-screen text-gray-100 px-6 py-12">
      <h1 className="text-3xl md:text-4xl font-bold text-center mb-12 text-[#25aae1] drop-shadow-[0_0_8px_#25aae1]">
        🛍️ Shop All T-Shirts
      </h1>

      {filteredProducts.length === 0 ? (
        <p className="text-center text-gray-400">No products found.</p>
      ) : (
        <div className="grid gap-10 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
          {filteredProducts.map((product) => (
            <div
              key={product.id}
              className="bg-[#121417] border border-[#2c2c2c] rounded-xl shadow-md hover:shadow-xl transition duration-300 flex flex-col justify-between"
            >
              <Link to={`/product/${product.slug}`}>
                <img
                  src={product.image}
                  alt={product.name}
                  className="w-full h-64 object-cover rounded-t-xl"
                />
              </Link>

              <div className="p-4 flex flex-col flex-grow justify-between">
                <div>
                  <h2 className="text-lg font-semibold mb-1 text-white line-clamp-2">
                    {product.name}
                  </h2>
                  <p className="text-sm text-[#25aae1] mb-2">
                    by {product.designer}
                  </p>
                  <p className="text-xl font-bold text-white">
                    ${product.price.toFixed(2)}
                  </p>
                  {product.discount && (
                    <p className="text-sm text-green-400 font-medium mt-1">
                      {product.discount}
                    </p>
                  )}
                </div>

                <Link to={`/product/${product.slug}`} className="mt-4">
                  <button className="w-full bg-[#25aae1] hover:bg-[#1f8fcb] text-black font-semibold py-2 rounded-lg transition duration-200">
                    View Product
                  </button>
                </Link>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default Shop;
