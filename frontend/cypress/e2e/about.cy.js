describe("About Page", () => {
    beforeEach(() => {
      cy.visit("/about"); // ⚠️ Adjust if route differs
    });
  
    it("renders the main quote", () => {
      cy.contains(/wearing a funny T-shirt is cheaper than therapy/i).should("be.visible");
    });
  
    it("displays the feature cards", () => {
      cy.contains(/premium quality/i).should("be.visible");
      cy.contains(/lightning delivery/i).should("be.visible");
      cy.contains(/eco friendly/i).should("be.visible");
    });
  
    it("shows the Tharun quote section", () => {
      cy.contains(/rey karthik.*custom print/i).should("be.visible");
      cy.contains(/not really tharun bhascker/i).should("be.visible");
    });
  
   
  });
  