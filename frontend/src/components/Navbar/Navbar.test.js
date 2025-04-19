import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import Navbar from "./Navbar";
import { Provider } from "react-redux";
import { BrowserRouter as Router } from "react-router-dom";
import configureStore from "redux-mock-store";

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

const renderComponent = () => {
  const store = mockStore(mockInitialState);
  render(
    <Provider store={store}>
      <Router>
        <Navbar />
      </Router>
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
    const badge = screen.getByText("4");
    expect(badge).toBeInTheDocument();
  });

  test("toggles mobile menu when menu button is clicked", () => {
    renderComponent();

    const menuButton = screen.getByRole("button", { name: /open menu/i });
    fireEvent.click(menuButton);

    expect(screen.getByRole("link", { name: /wishlist/i })).toBeInTheDocument();
    expect(screen.getByRole("link", { name: /cart/i })).toBeInTheDocument();
  });
});
