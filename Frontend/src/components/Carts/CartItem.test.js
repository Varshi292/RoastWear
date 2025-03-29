// CartItem.test.js
import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import CartItem from "./CartItem";
import { Provider } from "react-redux";
import configureStore from "redux-mock-store";
import { changeQuantity } from "../../stores/cart";

// Mock the product list (simulates ../Home/Product.js)
jest.mock("../Home/Product", () => ({
  products: [
    {
      id: 1,
      name: "Test T-Shirt",
      price: 10,
      image: "https://fakeimage.com/shirt.jpg",
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

    // Wait for useEffect to resolve
    const itemName = await screen.findByText("Test T-Shirt");
    const price = screen.getByText("$20");

    expect(itemName).toBeInTheDocument();
    expect(price).toBeInTheDocument();
    expect(screen.getByText("2")).toBeInTheDocument();
    expect(screen.getByRole("img")).toHaveAttribute("src", expect.stringContaining("shirt.jpg"));
  });

  test("clicking + calls dispatch with incremented quantity", async () => {
    setup();
    const plusButton = await screen.findByText("+");
    fireEvent.click(plusButton);

    expect(mockDispatch).toHaveBeenCalledWith(
      changeQuantity({ productId: 1, quantity: 3 })
    );
  });

  test("clicking - calls dispatch with decremented quantity", async () => {
    setup();
    const minusButton = await screen.findByText("-");
    fireEvent.click(minusButton);

    expect(mockDispatch).toHaveBeenCalledWith(
      changeQuantity({ productId: 1, quantity: 1 })
    );
  });
});
