describe("🛍 Shop Page", () => {
  beforeEach(() => {
    cy.visit("/shop");
  });

  it("renders the Shop page with title", () => {
    cy.contains("Shop All T-Shirts").should("be.visible");
  });

  it("renders products from mocked list", () => {
    cy.get('[data-testid="product-card"]').should("have.length.at.least", 1);
  });

  it("filters products using the search input", () => {
    cy.get('[data-testid="search-input"]').type("duck");
    cy.get('[data-testid="product-card"]').each(($el) => {
      cy.wrap($el).contains(/duck/i);
    });
  });

  it("navigates to product detail page", () => {
    cy.get('[data-testid="view-button"]').first().click();
    cy.url().should("include", "/product/");
  });
});
