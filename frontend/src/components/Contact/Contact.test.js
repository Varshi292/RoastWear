import React from "react";
import { render, screen } from "@testing-library/react";
import Contact from "./Contact";

describe("Contact Component", () => {
  beforeEach(() => {
    render(<Contact />);
  });



  it("renders section titles", () => {
    const sections = ["Shop", "About", "Help", "Social"];
    sections.forEach((title) => {
      expect(screen.getByText(title)).toBeInTheDocument();
    });
  });

  it("renders social icons with names", () => {
    const socials = ["Instagram", "Facebook", "Twitter", "Tumblr", "Pinterest"];
    socials.forEach((name) => {
      expect(screen.getByText(name)).toBeInTheDocument();
    });
  });

  it("renders footer legal links", () => {
    const legalLinks = [
      "User Agreement",
      "Privacy Policy",
      "Do not sell my personal information",
      "Cookie Policy",
    ];
    legalLinks.forEach((text) => {
      expect(screen.getByText(text)).toBeInTheDocument();
    });
  });

  it("renders copyright text", () => {
    expect(screen.getByText(/© RoastWear. All Rights Reserved/i)).toBeInTheDocument();
  });
});
