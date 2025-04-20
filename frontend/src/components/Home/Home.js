import React from "react";
import { Link } from "react-router-dom";
import { products } from "./Product";
import ProductCart from "../Carts/productCart";
import Contact from "../Contact/Contact.js";

const Home = () => {
  return (
    <div className="bg-[#0b0c0f] text-gray-600 min-h-screen">
      {/* Banner Section */}
      <div className="bg-[#25aae1] text-gray-600 p-8 text-center">
        <h1 className="text-4xl font-bold">Buy 2 @ $50</h1>
      </div>

      {/* Customization Banner */}
      <div className="bg-[#121417] p-8 text-center mt-8">
        <h2 className="text-3xl font-bold text-gray-300">Create Your Own T-Shirts!</h2>
        <p className="text-lg mt-4 text-[#5f6163]">
          Customize T-shirts with your favorite memes, upload your own images,
          or add personalized text!
        </p>
        <Link to="/customize">
          <button className="bg-[#22a7e0] text-gray-300 font-medium px-6 py-3 mt-6 rounded-lg hover:bg-[#1a9fd5] transition-all">
            Start Customizing
          </button>
        </Link>
      </div>

      {/* Categories Section */}
      <div className="mt-12 text-center">
        <h2 className="text-3xl font-medium text-[#ff2e63] drop-shadow-[0_0_8px_#ff2e63] mb-6">
          Browse Our Collections
        </h2>
        <div className="flex overflow-x-auto space-x-6 md:space-x-8 px-4 md:px-8 mb-10 scrollbar-hide">
          {[
            { to: "/shop/movie/hangover", src: "/Assets/hangover.webp", alt: "Hangover" },
            { to: "/shop/series/money-heist", src: "/Assets/moneyheist.jpg", alt: "Money Heist" },
            { to: "/shop/movie/bigbang", src: "/Assets/bigbang.avif", alt: "Big Bang Theory" },
            { to: "/shop/movie/brba", src: "/Assets/brba.webp", alt: "Breaking Bad" },
            { to: "/shop/movie/balayya", src: "/Assets/balayya.jpeg", alt: "balayya babu" },
            { to: "/shop/movie/friends", src: "/Assets/friends.avif", alt: "Friends" },
            { to: "/shop/movie/modernfamily", src: "/Assets/mrdnf.webp", alt: "modern family" },
            { to: "/shop/movie/ene", src: "/Assets/koushik.jpeg", alt: "ene" },
            { to: "/shop/movie/pellichoopulu", src: "/Assets/darshi.jpeg", alt: "pelli choopulu" }
          ].map(({ to, src, alt }, index) => (
            <Link key={index} to={to}>
              <img
                src={src}
                alt={alt}
                className="rounded-full object-cover cursor-pointer w-24 h-24 md:w-32 md:h-32 shadow-md hover:shadow-[0_0_12px_#ffffff50] transition duration-300"
              />
            </Link>
          ))}
        </div>
      </div>

      {/* Product Grid */}
      <div className="px-4 md:px-8">
        <h1 className="text-2xl md:text-3xl font-bold text-center text-gray-300 my-6">
          Welcome to Our T-Shirt Shop
        </h1>
        <div className="grid lg:grid-cols-4 md:grid-cols-3 sm:grid-cols-2 gap-5">
          {products.map((product, key) => (
            <ProductCart key={key} data={product} />
          ))}
        </div>
      </div>
      <div>

        
      </div>
      <div>
        <Contact />
      </div>
    </div>
  );
};

export default Home;
