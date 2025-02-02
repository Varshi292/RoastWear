import React from 'react';
import { Link } from 'react-router-dom';
import Product from './Product';

const Home = () => {
  return (
    <div>
      {/* Banner Section */}
      <div className="bg-blue-600 text-white p-8 text-center">
        <h1 className="text-4xl font-bold">Buy 2 @ $50</h1>
        <p className="text-xl mt-4">Premium Quality T-shirts with Free Shipping!</p>
      </div>

      {/* Customization Banner */}
      <div className="customization-banner bg-gray-100 p-8 text-center mt-8">
        <h2 className="text-3xl font-bold">Create Your Own T-Shirts!</h2>
        <p className="text-lg mt-4">
          Customize T-shirts with your favorite memes, upload your own images, or add personalized text!
        </p>
        <Link to="/customize">
          <button className="bg-blue-600 text-white px-6 py-3 mt-6 rounded-lg hover:bg-blue-700 transition-all">
            Start Customizing
          </button>
        </Link>
      </div>

      {/* Categories Section with Horizontal Scroll */}
      <div className="mt-12 text-center">
        <h2 className="text-3xl font-bold mb-6">Browse Our Collections</h2>
        <div className="flex overflow-x-auto space-x-6 md:space-x-8 px-4 md:px-8 mb-6 md:mb-10 scrollbar-hide">
  {/* Hangover */}
  <Link to="/shop/movie/hangover">
    <img
      src="/Assets/hangover.webp"
      alt="Hangover"
      className="rounded-full object-cover cursor-pointer w-24 h-24 sm:w-30 sm:h-30 md:w-32 md:h-32"
    />
  </Link>

  {/* Money Heist */}
  <Link to="/shop/series/money-heist">
    <img
      src="/Assets/moneyheist.jpg"
      alt="Money Heist"
      className="rounded-full object-cover cursor-pointer w-24 h-24 sm:w-30 sm:h-30 md:w-32 md:h-32"
    />
  </Link>

  {/* Big Bang */}
  <Link to="/shop/movie/bigbang">
    <img
      src="/Assets/bigbang.avif"
      alt="Big Bang Theory"
      className="rounded-full object-cover cursor-pointer w-24 h-24 sm:w-30 sm:h-30 md:w-32 md:h-32"
    />
  </Link>

  {/* Breaking Bad */}
  <Link to="/shop/movie/brba">
    <img
      src="/Assets/brba.webp"
      alt="Breaking Bad"
      className="rounded-full object-cover cursor-pointer w-24 h-24 sm:w-30 sm:h-30 md:w-32 md:h-32"
    />
  </Link>

  {/* Friends */}
  <Link to="/shop/movie/friends">
    <img
      src="/Assets/friends.avif"
      alt="Friends"
      className="rounded-full object-cover cursor-pointer w-24 h-24 sm:w-30 sm:h-30 md:w-32 md:h-32"
    />
  </Link>
</div>

      </div>
       
      {/* Product Component */}
      <div className="home-container px-4 md:px-8">
        <h1 className="text-2xl md:text-3xl font-bold text-center my-6">
          Welcome to Our T-Shirt Shop
        </h1>
        <Product />
      </div>
    </div>
  );
};

export default Home;
