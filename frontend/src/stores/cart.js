// stores/cart.js
import { createSlice } from "@reduxjs/toolkit";

// LocalStorage sync on load
const initialState = {
  items: localStorage.getItem("carts")
    ? JSON.parse(localStorage.getItem("carts"))
    : [],
  statusTab: false,
};

const cartSlice = createSlice({
  name: "cart",
  initialState,
  reducers: {
    // ✅ Add or increment product in cart
    addToCart(state, action) {
      const { productId, quantity } = action.payload;
      const index = state.items.findIndex(
        (item) => item.productId === productId
      );

      if (index >= 0) {
        state.items[index].quantity += quantity;
      } else {
        state.items.push({ productId, quantity });
      }

      localStorage.setItem("carts", JSON.stringify(state.items));
    },

    // ✅ Change quantity or remove if zero
    changeQuantity(state, action) {
      const { productId, quantity } = action.payload;
      const index = state.items.findIndex(
        (item) => item.productId === productId
      );

      if (quantity > 0 && index >= 0) {
        state.items[index].quantity = quantity;
      } else {
        state.items = state.items.filter(
          (item) => item.productId !== productId
        );
      }

      localStorage.setItem("carts", JSON.stringify(state.items));
    },

    // ✅ NEW: Remove item entirely
    removeFromCart(state, action) {
      const productId = action.payload;
      state.items = state.items.filter((item) => item.productId !== productId);
      localStorage.setItem("carts", JSON.stringify(state.items));
    },

    // ✅ Toggle cart panel visibility
    toggleStatusTab(state) {
      state.statusTab = !state.statusTab;
    },
  },
});

export const {
  addToCart,
  changeQuantity,
  removeFromCart, // <- Make sure to export
  toggleStatusTab,
} = cartSlice.actions;

export default cartSlice.reducer;
