describe("Customize Page", () => {
    beforeEach(() => {
      cy.visit("/customize");
    });
  
    it("renders the page correctly", () => {
      cy.get("[data-testid=customize-root]").should("exist");
      cy.contains("Customize Your T-Shirts").should("be.visible");
    });
  
    it("changes shirt color when buttons are clicked", () => {
      const colors = ["white", "black", "red", "blue"];
  
      colors.forEach((color) => {
        cy.get(`button[aria-label="Select ${color} shirt"]`).click();
      });
    });
  
    it("accepts overlay text", () => {
      cy.get("#overlay-text").type("Cypress Rocks!");
      cy.get("#overlay-text").should("have.value", "Cypress Rocks!");
    });
  
    it("updates text color with color input", () => {
      cy.get("#text-color").invoke("val", "#ff0000").trigger("input");
      cy.get("#text-color").should("have.value", "#ff0000");
    });
  
    it("handles Tenor GIF input", () => {
      cy.get("#tenor-url")
        .type("https://media.tenor.com/Ljz5GXdp3kQAAAAC/hi.gif")
        .should("have.value", "https://media.tenor.com/Ljz5GXdp3kQAAAAC/hi.gif");
    });
  
    
  });
  