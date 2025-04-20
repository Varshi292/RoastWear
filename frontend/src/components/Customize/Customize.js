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
    <div data-testid="customize-root">
      <h2>Customize Your T-Shirts</h2>

      {/* Shirt Color Buttons */}
      <div>
        {["white", "black", "red", "blue"].map((color) => (
          <button
            key={color}
            onClick={() => setShirtColor(color)}
            aria-label={`Select ${color} shirt`}
          >
            {color}
          </button>
        ))}
      </div>

      {/* Shirt Preview */}
      <div>
        <img src={getShirtImage()} alt="T-shirt mockup" />
        {previewImg && (
          <Draggable bounds="parent">
            <img src={previewImg} alt="Preview" />
          </Draggable>
        )}
        {overlayText && (
          <Draggable bounds="parent">
            <div style={{ color: textColor }}>{overlayText}</div>
          </Draggable>
        )}
      </div>

      {/* Inputs */}
      <div>
        <label htmlFor="tenor-url">Tenor GIF URL:</label>
        <input
          id="tenor-url"
          placeholder="https://media.tenor.com/xyz.gif"
          value={tenorLink}
          onChange={handleTenorLinkChange}
        />

        <label htmlFor="upload-image">Upload Image:</label>
        <input
          id="upload-image"
          type="file"
          accept="image/*"
          onChange={handleFileChange}
        />

        <label htmlFor="overlay-text">Overlay Text:</label>
        <input
          id="overlay-text"
          placeholder="e.g. Meme Lords Only"
          value={overlayText}
          onChange={handleTextChange}
        />

        <label htmlFor="text-color">Text Color:</label>
        <input
          id="text-color"
          type="color"
          value={textColor}
          onChange={(e) => setTextColor(e.target.value)}
        />
      </div>
    </div>
  );
};

export default Customize;
