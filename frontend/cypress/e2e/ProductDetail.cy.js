describe('Product Detail Page', () => {
  beforeEach(() => {
    cy.visit('http://localhost:3000/product/tobe-fonseca'); 
    cy.wait(1000); 
  });

  it('displays product information', () => {
    cy.contains('I Choose Violence Funny Duck T-Shirt').should('exist');
    cy.contains('$19.25').should('exist');
    cy.get('img[alt="I Choose Violence Funny Duck T-Shirt"]').should('be.visible');
  });
  

  it('updates quantity', () => {
    cy.get('button').contains('+').click();
    cy.contains('2').should('exist');
    cy.get('button').contains('−').click(); // Uses proper minus symbol!
    cy.contains('1').should('exist');
  });

  it('adds product to cart', () => {
    cy.get('button').contains('Add to Cart').click();
    // No cart assertion here unless UI shows confirmation
  });
});