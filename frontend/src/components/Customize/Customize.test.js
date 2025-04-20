import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import Customize from "./Customize";

jest.mock("react-draggable", () => ({
  __esModule: true,
  default: ({ children }) => <div>{children}</div>,
}));


describe("Customize Component", () => {
  it("renders title and color buttons", () => {
    render(<Customize />);
    expect(screen.getByText(/Customize Your T-Shirts/i)).toBeInTheDocument();
    expect(screen.getAllByRole("button")).toHaveLength(4);
  });

  it("shows overlay text", () => {
    render(<Customize />);
    fireEvent.change(screen.getByPlaceholderText(/meme/i), {
      target: { value: "🔥 test" },
    });
    expect(screen.getByText("🔥 test")).toBeInTheDocument();
  });


});
