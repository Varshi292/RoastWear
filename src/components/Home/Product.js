import React from "react";

const productsData = [
  {
    id: 1,
    name: "I Choose Violence Funny Duck T-Shirt",
    designer: "Tobe Fonseca",
    price: 19.25,
    discount: "25% off",
    image:
      "https://ih1.redbubble.net/image.10557653.9817/ssrco,slim_fit_t_shirt,womens,fafafa:ca443f4786,front,square_product,600x600.u7.jpg",
  },
  {
    id: 2,
    name: "Another Cool T-Shirt",
    designer: "Designer Name",
    price: 24.99,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1600781373074-529b34723d11?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDE0fHx0LXNoaXJ0fGVufDB8fHx8fDE2MzU3MjM0Nzk&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 3,
    name: "Graphic Art Tee",
    designer: "Jane Doe",
    price: 29.99,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1600264519827-dfbf9a83f9db?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDJ8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 4,
    name: "Retro Sunset T-Shirt",
    designer: "Alex Lee",
    price: 19.99,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1571774120840-86e05a8b3e16?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDF8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 5,
    name: "Bold Street Art Tee",
    designer: "Chris Wright",
    price: 22.99,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1609454173522-04fa0523a9c4?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDJ8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 6,
    name: "Urban Explorer T-Shirt",
    designer: "Michael Kingsley",
    price: 27.99,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1593642634311-d6f3c1c1d4c9?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDJ8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 7,
    name: "Space Adventure T-Shirt",
    designer: "Sarah Turner",
    price: 21.5,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1603996849377-47a8b47891e0?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDd8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 8,
    name: "Minimalist Tee",
    designer: "Sophia Adams",
    price: 18.5,
    discount: "15% off",
    image:
      "https://images.unsplash.com/photo-1629380477757-20fd315e1261?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDg4fHx0LXNoaXJ0fGVufDB8fHx8fDE2NzM4NzU0Mzc&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 9,
    name: "Abstract Art Tee",
    designer: "Emily Robinson",
    price: 20.99,
    discount: "10% off",
    image:
      "https://images.unsplash.com/photo-1593784993572-0dded5c03029?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDEwfHx0LXNoaXJ0fGVufDB8fHx8fDE2NzM4NzU0Mzc&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 10,
    name: "Vintage Travel Tee",
    designer: "Olivia Johnson",
    price: 23.99,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1512436991641-6745cdb1723f?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDJ8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 11,
    name: "Eco-friendly Tee",
    designer: "Ella Brooks",
    price: 25.99,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1617634462586-ef5d5479f9be?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDN8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 12,
    name: "Monochrome Tee",
    designer: "Benjamin Taylor",
    price: 22.49,
    discount: "5% off",
    image:
      "https://images.unsplash.com/photo-1598454449594-1b1c02b89e9e?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDN8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 13,
    name: "Ocean Waves Tee",
    designer: "Nathan Parker",
    price: 19.75,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1600992450680-13fa6c60adcd?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDZ8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 14,
    name: "Mountain Explorer Tee",
    designer: "Hannah Scott",
    price: 28.99,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1599460720587-5683d361f1db?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDZ8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 15,
    name: "City Skyline Tee",
    designer: "Daniel Green",
    price: 24.5,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1527600978858-529073c33422?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDE2fHx0LXNoaXJ0fGVufDB8fHx8fDE2NzM4NzU0Mzc&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 16,
    name: "Modern Abstract Tee",
    designer: "Zara Williams",
    price: 26.5,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1580658337840-13f8f7ff4b64?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDIwfHx0LXNoaXJ0fGVufDB8fHx8fDE2NzM4NzU0Mzc&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 17,
    name: "Monochrome Tee",
    designer: "Benjamin Taylor",
    price: 22.49,
    discount: "5% off",
    image:
      "https://images.unsplash.com/photo-1598454449594-1b1c02b89e9e?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDN8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
  {
    id: 18,
    name: "Ocean Waves Tee",
    designer: "Nathan Parker",
    price: 19.75,
    discount: null,
    image:
      "https://images.unsplash.com/photo-1600992450680-13fa6c60adcd?crop=entropy&cs=tinysrgb&fit=max&ixid=MnwzNjk2OXwwfDF8c2VhcmNofDZ8fHRzaGlydHxlbnwwfDB8fHx8&ixlib=rb-1.2.1&q=80&w=400",
  },
];

const Products = () => {
  return (
    <div className="products-container grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-6 p-6">
      {productsData.map((product) => (
        <div
          key={product.id}
          className="product group bg-white shadow-lg rounded-lg overflow-hidden relative"
        >
          {/* Add to Wishlist Icon */}

          {/* Product Image */}
          <div className="relative overflow-hidden">
            <img
              src={product.image}
              alt={product.name}
              className="product-image w-full h-64 object-cover transform transition duration-300 group-hover:scale-110"
            />
          </div>
          <div className="absolute top-2 right-2">
            <button className="wishlist-btn bg-white p-2 rounded-full shadow-md hover:bg-red-100 transition duration-300">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
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

          {/* Product Info */}
          <div className="product-info p-4">
  {/* Product Name */}
  <h3 className="text-lg sm:text-xl md:text-2xl font-semibold truncate">{product.name}</h3>

  {/* Designer Info */}
  <p className="text-gray-600 text-xs sm:text-sm">{product.designer}</p>

  {/* Price */}
  <p className="price text-base sm:text-lg md:text-xl text-red-500">
    ${product.price.toFixed(2)}{" "}
    {product.discount && (
      <span className="text-gray-500 text-xs sm:text-sm">
        ({product.discount} off)
      </span>
    )}
  </p>

  {/* Add to Cart Button */}
  <button className="add-to-cart mt-4 bg-blue-500 text-white py-1.5 px-4 sm:py-2 sm:px-5 md:py-2.5 md:px-6 rounded-full hover:bg-blue-600 transition duration-300 text-xs sm:text-sm">
    Add to Cart
  </button>
</div>


        </div>
      ))}
    </div>
  );
};

export default Products;