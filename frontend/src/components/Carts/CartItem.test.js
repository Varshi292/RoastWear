import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import CartItem from "./CartItem";
import { Provider } from "react-redux";
import configureStore from "redux-mock-store";
import { changeQuantity, removeFromCart } from "../../stores/cart";

// Mock product data
jest.mock("../Home/Product", () => ({
  products: [
    {
      id: 1,
      name: "Test T-Shirt",
      price: 10,
      image: "https://fakeimage.com/shirt.jpg",
      designer: "Test Designer",
    },
  ],
}));

const mockStore = configureStore([]);
const mockDispatch = jest.fn();

jest.mock("react-redux", () => ({
  ...jest.requireActual("react-redux"),
  useDispatch: () => mockDispatch,
}));

describe("CartItem Component", () => {
  let store;

  beforeEach(() => {
    store = mockStore({});
    mockDispatch.mockClear();
  });

  const setup = () => {
    const props = {
      data: {
        productId: 1,
        quantity: 2,
      },
    };

    render(
      <Provider store={store}>
        <CartItem {...props} />
      </Provider>
    );
  };

  test("renders item details correctly", async () => {
    setup();
    const itemName = await screen.findByText("Test T-Shirt");
    const price = screen.getByText((content) => content.includes("20.00"));

    expect(itemName).toBeInTheDocument();
    expect(price).toBeInTheDocument();
    expect(screen.getByText("2")).toBeInTheDocument();
    expect(screen.getByRole("img")).toHaveAttribute(
      "src",
      expect.stringContaining("shirt.jpg")
    );
  });

  test("clicking + calls dispatch with incremented quantity", async () => {
    setup();
    const plusButton = screen.getByRole("button", { name: /plus/i });
    fireEvent.click(plusButton);

    expect(mockDispatch).toHaveBeenCalledWith(
      changeQuantity({ productId: 1, quantity: 3 })
    );
  });

  test("clicking - calls dispatch with decremented quantity", async () => {
    setup();
    const minusButton = screen.getByRole("button", { name: /minus/i });
    fireEvent.click(minusButton);

    expect(mockDispatch).toHaveBeenCalledWith(
      changeQuantity({ productId: 1, quantity: 1 })
    );
  });

  test("clicking trash calls removeFromCart", async () => {
    setup();
    const removeButton = screen.getByRole("button", { name: /remove item/i });
    fireEvent.click(removeButton);

    expect(mockDispatch).toHaveBeenCalledWith(removeFromCart(1));
  });
});
