/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        roast: {
          bg: '#f3efe9',      // Light cream background
          dark: '#0b0c0f',    // Deep dark like cinema hall
          neon: '#ff2851',    // Bright neon pink
          blue: '#25aae1',    // Chill aqua blue
          mustard: '#ffd369', // Fun punchy yellow
          text: '#2f2f2f',    // Main text color
          sub: '#9ca3af',     // Subtle muted text
          border: '#e5e7eb',  // Light borders
        },
      },
      fontFamily: {
        sans: ['"Inter"', 'sans-serif'],
      },
    },
  },
  plugins: [],
};
