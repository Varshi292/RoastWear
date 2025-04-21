import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { Provider } from "react-redux";
import { store, persistor } from "./stores/store";
import { PersistGate } from "redux-persist/integration/react";
import { BrowserRouter } from "react-router-dom";
import { UserProvider } from "./components/Context/UserContext";
import { SearchProvider } from "./components/Context/SearchContext";

const root = ReactDOM.createRoot(document.getElementById("root"));

root.render(
  <React.StrictMode>
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <BrowserRouter>
          <SearchProvider>
            <UserProvider>
              <App />
            </UserProvider>
          </SearchProvider>
        </BrowserRouter>
      </PersistGate>
    </Provider>
  </React.StrictMode>
);