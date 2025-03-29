// CartTab.test.js
import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import CartTab from "./CartTab";
import { Provider } from "react-redux";
import configureStore from "redux-mock-store";

const mockDispatch = jest.fn();

// 🧪 Fixed Mock for react-redux
jest.mock("react-redux", () => {
  const ActualRedux = jest.requireActual("react-redux");
  return {
    ...ActualRedux,
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

  test("renders title and items", () => {
    renderWithStore();
    expect(screen.getByText("Shopping Cart")).toBeInTheDocument();
    expect(screen.getAllByTestId("cart-item")).toHaveLength(2);
  });

  test("clicking CLOSE dispatches toggle action", () => {
    renderWithStore();
    fireEvent.click(screen.getByText("CLOSE"));
    expect(mockDispatch).toHaveBeenCalled();
  });

  test("hides component if statusTab is false", () => {
    store = mockStore({
      cart: { items: [], statusTab: false },
    });
    render(
      <Provider store={store}>
        <CartTab />
      </Provider>
    );
    const tab = screen.getByText("Shopping Cart").parentElement;
    expect(tab.className).toContain("translate-x-full");
  });
});
