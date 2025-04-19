// src/stores/wishlist.js
import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  items: [], // productId array
};

const wishlistSlice = createSlice({
  name: "wishlist",
  initialState,
  reducers: {
    toggleWishlist(state, action) {
      const productId = action.payload;
      if (state.items.includes(productId)) {
        state.items = state.items.filter((id) => id !== productId);
      } else {
        state.items.push(productId);
      }
    },
  },
});

export const { toggleWishlist } = wishlistSlice.actions;
export default wishlistSlice.reducer;
