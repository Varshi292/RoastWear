module.exports = {
  testEnvironment: "jsdom",
  transform: {
    "^.+\\.[jt]sx?$": "babel-jest",
  },
  transformIgnorePatterns: [
    "/node_modules/(?!(swiper|ssr-window|dom7)/)",
  ],
  moduleNameMapper: {
    "^swiper/react$": "<rootDir>/__mocks__/swiper/react.js",
  },
};
