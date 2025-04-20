import React from "react";
import { render, screen } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import Shop from "./Shop";
import { SearchProvider } from "../Context/SearchContext"; // ✅ Required for search context

// 🧪 Mock product list
jest.mock("../Home/Product", () => ({
  products: [
    {
      id: 1,
      name: "Funny Duck T-Shirt",
      designer: "Tobe Fonseca",
      price: 19.25,
      discount: "25% off",
      slug: "duck-tee",
      image: "mock-image.jpg",
    },
    {
      id: 2,
      name: "Abstract Tee",
      designer: "Zara Williams",
      price: 21.99,
      discount: null,
      slug: "abstract-tee",
      image: "mock-image.jpg",
    },
  ],
}));

describe("🧪 Shop Component", () => {
  beforeEach(() => {
    render(
      <SearchProvider>
        <BrowserRouter>
          <Shop />
        </BrowserRouter>
      </SearchProvider>
    );
  });

  it("renders the shop heading", () => {
    expect(screen.getByText(/Shop All T-Shirts/i)).toBeInTheDocument();
  });

  it("displays all mocked products", () => {
    expect(screen.getByText(/Funny Duck T-Shirt/i)).toBeInTheDocument();
    expect(screen.getByText(/Abstract Tee/i)).toBeInTheDocument();
  });

  it("renders View Product buttons", () => {
    const buttons = screen.getAllByRole("button", { name: /View Product/i });
    expect(buttons.length).toBe(2);
  });
});
