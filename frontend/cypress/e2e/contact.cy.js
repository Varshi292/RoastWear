describe("Contact Component", () => {
    beforeEach(() => {
      cy.visit("/contact"); 
    });
  
    it("renders Contact Us heading", () => {
      cy.contains("Contact Us").should("be.visible");
    });
  
    it("renders all 4 section headers", () => {
      const sections = ["Shop", "About", "Help", "Social"];
      sections.forEach((title) => {
        cy.contains(title).should("be.visible");
      });
    });
  
    it("renders social media icons and names", () => {
      const socialItems = ["Instagram", "Facebook", "Twitter", "Tumblr", "Pinterest"];
      socialItems.forEach((name) => {
        cy.contains(name).should("be.visible");
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
        cy.contains(text).should("be.visible");
      });
    });
  
    it("renders copyright", () => {
      cy.contains("© RoastWear. All Rights Reserved").should("be.visible");
    });
  });
  