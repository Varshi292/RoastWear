import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import CartTab from "./CartTab";
import { Provider } from "react-redux";
import configureStore from "redux-mock-store";

const mockDispatch = jest.fn();

jest.mock("react-redux", () => {
  const actual = jest.requireActual("react-redux");
  return {
    ...actual,
    useDispatch: () => mockDispatch,
  };
});

jest.mock("./CartItem", () => (props) => (
  <div data-testid="cart-item">{JSON.stringify(props.data)}</div>
));

const mockStore = configureStore([]);

describe("CartTab Component", () => {
  let store;

  beforeEach(() => {
    store = mockStore({
      cart: {
        items: [
          { productId: 1, quantity: 2 },
          { productId: 2, quantity: 1 },
        ],
        statusTab: true,
      },
    });
    mockDispatch.mockClear();
  });

  const renderWithStore = () =>
    render(
      <Provider store={store}>
        <CartTab />
      </Provider>
    );

  test("renders title and items when open", () => {
    renderWithStore();
    expect(screen.getByText("🛒 Your Cart")).toBeInTheDocument();
    expect(screen.getAllByTestId("cart-item")).toHaveLength(2);
  });

  test("clicking ✕ button dispatches toggleStatusTab", () => {
    renderWithStore();
    const closeBtn = screen.getByRole("button", { name: /close cart/i });
    fireEvent.click(closeBtn);
    expect(mockDispatch).toHaveBeenCalled();
  });

  test("does not render if statusTab is false", () => {
    store = mockStore({
      cart: { items: [], statusTab: false },
    });
    render(
      <Provider store={store}>
        <CartTab />
      </Provider>
    );
    expect(screen.queryByText("🛒 Your Cart")).not.toBeInTheDocument();
  });
});
