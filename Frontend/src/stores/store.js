import { configureStore } from '@reduxjs/toolkit';
import cartReducer from './cart';
import wishlistReducer from './wishlist'; // 👈 important

const store = configureStore({
  reducer: {
    cart: cartReducer,
    wishlist: wishlistReducer, // 👈 important
  },
});

export default store;

