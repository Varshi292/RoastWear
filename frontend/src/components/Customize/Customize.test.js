import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import Customize from "./Customize";

// Mock draggable for simplicity in tests
jest.mock("react-draggable", () => ({
  __esModule: true,
  default: ({ children }) => <div>{children}</div>,
}));

describe("Customize Component", () => {
  it("renders title and 4 shirt color buttons", () => {
    render(<Customize />);
    expect(screen.getByText(/Customize Your T-Shirts/i)).toBeInTheDocument();

    const colorButtons = screen.getAllByRole("button", { name: /select/i });
    expect(colorButtons).toHaveLength(4);
  });

  it("shows overlay text when typed", () => {
    render(<Customize />);
    fireEvent.change(screen.getByPlaceholderText(/meme/i), {
      target: { value: "🔥 test" },
    });
    expect(screen.getByText("🔥 test")).toBeInTheDocument();
  });
});
