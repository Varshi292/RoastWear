import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { BrowserRouter } from "react-router-dom";
import { Provider } from "react-redux";
import store from "./stores/store";
import { SearchProvider } from "./context/SearchContext";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
// ✅ index.js
<Provider store={store}>
  <PersistGate loading={null} persistor={persistor}>
    <BrowserRouter> 
      <SearchProvider>
        <App />
      </SearchProvider>
    </BrowserRouter>
  </PersistGate>
</Provider>

  </React.StrictMode>
);
