import { createSlice } from "@reduxjs/toolkit";

// Backend sync helper
export const modifyCartBackend = async ({ username, sessionid, productId, quantity, price }) => {
  try {
    const response = await fetch("http://localhost:7777/cart/modify", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username,
        sessionid: sessionid || "placeholder",
        productid: productId,
        quantity,
        unitPrice: price * quantity,
      }),
    });

    const result = await response.json();
    if (result.message) {
      console.log("✅", result.message);
    }
  } catch (err) {
    console.error("❌ Failed to modify cart:", err);
  }
};

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
    addToCart(state, action) {
      const { productId, quantity, price, username, sessionid } = action.payload;
      const index = state.items.findIndex((item) => item.productId === productId);

      if (index >= 0) {
        state.items[index].quantity += quantity;
      } else {
        state.items.push({ productId, quantity });
      }

      localStorage.setItem("carts", JSON.stringify(state.items));

      const updatedQuantity = state.items.find(i => i.productId === productId)?.quantity || quantity;
      modifyCartBackend({ username, sessionid, productId, quantity: updatedQuantity, price });
    },

    changeQuantity(state, action) {
      const { productId, quantity, price, username, sessionid } = action.payload;
      const index = state.items.findIndex((item) => item.productId === productId);

      if (quantity > 0 && index >= 0) {
        state.items[index].quantity = quantity;
      } else {
        state.items = state.items.filter((item) => item.productId !== productId);
      }

      localStorage.setItem("carts", JSON.stringify(state.items));
      modifyCartBackend({ username, sessionid, productId, quantity, price });
    },

    removeFromCart(state, action) {
      const { productId, username, sessionid, price } = action.payload;
      state.items = state.items.filter((item) => item.productId !== productId);
      localStorage.setItem("carts", JSON.stringify(state.items));
      modifyCartBackend({ username, sessionid, productId, quantity: 0, price });
    },

    clearCart(state) {
      state.items = [];
      localStorage.removeItem("carts");
    },    

    toggleStatusTab(state) {
      state.statusTab = !state.statusTab;
    },
  },
});

export const {
  addToCart,
  changeQuantity,
  removeFromCart,
  toggleStatusTab,
  clearCart,
} = cartSlice.actions;

export default cartSlice.reducer;
