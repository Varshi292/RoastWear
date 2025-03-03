import React, { useEffect } from 'react';

const Cart = () => {
  useEffect(() => {
    const scriptId = "tenor-embed-script";
    
    // Check if the script is already present
    if (!document.getElementById(scriptId)) {
      const script = document.createElement("script");
      script.id = scriptId;
      script.src = "https://tenor.com/embed.js";
      script.async = true;
      document.body.appendChild(script);
    }

    return () => {
      // Clean up script on unmount to prevent conflicts
      const existingScript = document.getElementById(scriptId);
      if (existingScript) {
        existingScript.remove();
      }
    };
  }, []);

  return (
    <div className="flex justify-center items-center h-screen rounded-md">
      <div className="w-[500px] h-[500px] overflow-hidden">
        <div
          className="tenor-gif-embed w-full h-full"
          data-postid="19048359"
          data-share-method="host"
          data-aspect-ratio="1.40969"
          data-width="100%"
        >
          <a href="https://tenor.com/view/sad-emotional-pain-venkatesh-gif-gif-19048359">
            Sad Emotional GIF
          </a>
        </div>
      </div>
    </div>
  );
};

export default Cart;