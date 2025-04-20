import React from "react";
import { render, screen } from "@testing-library/react";
import Navbar from "./Navbar";
import { Provider } from "react-redux";
import { BrowserRouter as Router } from "react-router-dom";
import configureStore from "redux-mock-store";

// ✅ Import your SearchContext
import { SearchContext } from "../Context/SearchContext";

const mockStore = configureStore([]);

const mockInitialState = {
  cart: {
    items: [
      { productId: 1, quantity: 1 },
      { productId: 2, quantity: 3 },
    ],
    statusTab: false,
  },
};

// ✅ Utility to render Navbar with all necessary providers
const renderComponent = () => {
  const store = mockStore(mockInitialState);
  const mockSearchContext = {
    searchTerm: "",
    setSearchTerm: jest.fn(),
  };

  render(
    <Provider store={store}>
      <SearchContext.Provider value={mockSearchContext}>
        <Router>
          <Navbar />
        </Router>
      </SearchContext.Provider>
    </Provider>
  );
};

describe("Navbar Component", () => {
  test("renders navigation links", () => {
    renderComponent();
    expect(screen.getAllByText(/home/i).length).toBeGreaterThan(0);
    expect(screen.getAllByText(/shop/i).length).toBeGreaterThan(0);
    expect(screen.getAllByText(/customize/i).length).toBeGreaterThan(0);
    expect(screen.getAllByText(/about/i).length).toBeGreaterThan(0);
    expect(screen.getAllByText(/login/i).length).toBeGreaterThan(0);
  });

  test("renders cart quantity badge", () => {
    renderComponent();
    expect(screen.getByText("4")).toBeInTheDocument(); // 1 + 3
  });

});
