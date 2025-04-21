// src/components/Banners/WelcomeBanner.js
import React from "react";
import { useUser } from "../Context/UserContext";

const WelcomeBanner = () => {
  const { userName } = useUser();

  if (!userName) return null;

  return (
    <div className="bg-[#121417] text-white py-6 px-4 text-center text-2xl font-semibold">
      Welcome{" "}
      <span className="text-[#ff2e63] text-3xl font-bold drop-shadow">
        {userName}
      </span>{" "}
      – Wanna make your closet the funniest place on earth! 
    </div>
  );
};

export default WelcomeBanner;
