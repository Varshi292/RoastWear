// Home.test.js
import React from "react";
import { render, screen } from "@testing-library/react";
import Home from "./Home";
import { MemoryRouter } from "react-router-dom";

// Mock products
jest.mock("./Product", () => ({
  products: [
    { id: 1, title: "T-Shirt One", price: 25 },
    { id: 2, title: "T-Shirt Two", price: 30 },
  ],
}));

// Mock ProductCart component
jest.mock("../Carts/productCart", () => ({ data }) => (
  <div data-testid="product-cart">{data.title}</div>
));

// Mock Contact component
jest.mock("../Contact.js", () => () => <div data-testid="contact-section">Contact Component</div>);

describe("Home Component", () => {
  beforeEach(() => {
    render(
      <MemoryRouter>
        <Home />
      </MemoryRouter>
    );
  });

  test("renders banner section", () => {
    expect(screen.getByText(/buy 2 @ \$50/i)).toBeInTheDocument();
    expect(screen.getByText(/premium quality t-shirts/i)).toBeInTheDocument();
  });

  test("renders customization banner with button", () => {
    expect(screen.getByText(/create your own t-shirts/i)).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /start customizing/i })).toBeInTheDocument();
  });

  test("renders category image links", () => {
    const altTags = [
      "Hangover",
      "Money Heist",
      "Big Bang Theory",
      "Breaking Bad",
      "balayya babu",
      "Friends",
      "modern family",
      "ene",
      "pelli choopulu",
    ];

    altTags.forEach((altText) => {
      expect(screen.getByAltText(altText)).toBeInTheDocument();
    });
  });

  test("renders products from mocked data", () => {
    const products = screen.getAllByTestId("product-cart");
    expect(products).toHaveLength(2);
    expect(products[0]).toHaveTextContent("T-Shirt One");
    expect(products[1]).toHaveTextContent("T-Shirt Two");
  });

  test("renders contact component", () => {
    expect(screen.getByTestId("contact-section")).toBeInTheDocument();
  });
});
