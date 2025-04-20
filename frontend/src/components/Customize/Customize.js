import React, { useState } from "react";
import Draggable from "react-draggable";
import { convertGifToStatic } from "../../utils/convertGifToStatic";

// Shirt Assets
import whiteShirt from "../../assets/images/tshirt-white.jpg";
import blackShirt from "../../assets/images/tshirt-black.webp";
import redShirt from "../../assets/images/tshirt-red.webp";
import blueShirt from "../../assets/images/tshirt-blue.webp";

const Customize = () => {
  const [tenorLink, setTenorLink] = useState("");
  const [uploadedImg, setUploadedImg] = useState(null);
  const [overlayText, setOverlayText] = useState("");
  const [staticGifImage, setStaticGifImage] = useState(null);
  const [textColor, setTextColor] = useState("#ffffff");
  const [shirtColor, setShirtColor] = useState("white");

  const getShirtImage = () => {
    switch (shirtColor) {
      case "black": return blackShirt;
      case "red": return redShirt;
      case "blue": return blueShirt;
      default: return whiteShirt;
    }
  };

  const handleTenorLinkChange = async (e) => {
    const url = e.target.value;
    setUploadedImg(null);
    setOverlayText("");
    setTenorLink(url);
    setStaticGifImage(null);

    if (!url) return;

    try {
      const staticImg = await convertGifToStatic(url);
      setStaticGifImage(staticImg);
    } catch (err) {
      console.error("GIF conversion failed:", err);
    }
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = () => {
      setUploadedImg(reader.result);
      setTenorLink("");
      setOverlayText("");
      setStaticGifImage(null);
    };
    reader.readAsDataURL(file);
  };

  const handleTextChange = (e) => {
    setOverlayText(e.target.value);
    setTenorLink("");
    setUploadedImg(null);
    setStaticGifImage(null);
  };

  const previewImg = staticGifImage || uploadedImg;

  return (
    <div className="min-h-screen bg-[#0b0c0f] text-gray-300 px-6 py-10 space-y-10">
      <h2 className="text-3xl font-bold text-center text-[#25aae1] drop-shadow-[0_0_8px_#25aae1]">
        Customize Your T-Shirts
      </h2>

      {/* Shirt Color Picker */}
      <div className="flex justify-center space-x-4 mb-6">
        {["white", "black", "red", "blue"].map((color) => (
          <button
            key={color}
            onClick={() => setShirtColor(color)}
            className="w-8 h-8 rounded-full border border-gray-600"
            style={{ backgroundColor: color }}
            aria-label={`Select ${color} shirt`}
          ></button>
        ))}
      </div>

      {/* T-Shirt Preview */}
      <div className="flex justify-center">
        <div className="relative w-[300px] h-[400px] bg-[#1f1f1f] shadow-xl rounded-lg overflow-hidden">
          <img
            src={getShirtImage()}
            alt="T-shirt mockup"
            className="w-full h-full object-contain"
          />

          {/* Image */}
          {previewImg && (
            <Draggable bounds="parent">
              <img
                src={previewImg}
                alt="Preview"
                className="absolute top-[30%] left-1/2 w-[150px] max-h-[120px] -translate-x-1/2 object-contain cursor-move"
              />
            </Draggable>
          )}

          {/* Text */}
          {overlayText && (
            <Draggable bounds="parent">
              <div
                className="absolute top-[60%] left-1/2 -translate-x-1/2 text-center bg-black/50 px-3 py-1 rounded text-sm font-semibold cursor-move"
                style={{ color: textColor }}
              >
                {overlayText}
              </div>
            </Draggable>
          )}
        </div>
      </div>

      {/* Input Fields */}
      <div className="max-w-2xl mx-auto space-y-6">
        <div>
          <label className="font-semibold block mb-1">Tenor GIF URL:</label>
          <input
            type="text"
            value={tenorLink}
            onChange={handleTenorLinkChange}
            className="w-full p-3 border rounded-md bg-[#1a1a1a] text-white placeholder:text-gray-400"
            placeholder="https://media.tenor.com/xyz.gif"
          />
        </div>

        <div>
          <label className="font-semibold block mb-1">Upload Image:</label>
          <input
            type="file"
            accept="image/*"
            onChange={handleFileChange}
            className="w-full text-white"
          />
        </div>

        <div>
          <label className="font-semibold block mb-1">Overlay Text:</label>
          <input
            type="text"
            value={overlayText}
            onChange={handleTextChange}
            className="w-full p-3 border rounded-md bg-[#1a1a1a] text-white placeholder:text-gray-400"
            placeholder="e.g. Meme Lords Only"
          />
        </div>

        <div>
          <label className="font-semibold block mb-1">Text Color:</label>
          <input
            type="color"
            value={textColor}
            onChange={(e) => setTextColor(e.target.value)}
            className="w-12 h-10 p-1 border rounded-md bg-[#1a1a1a]"
          />
        </div>
      </div>
    </div>
  );
};

export default Customize;
