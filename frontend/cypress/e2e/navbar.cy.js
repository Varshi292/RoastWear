describe('Navbar Component', () => {
    beforeEach(() => {
      cy.visit('/');
    });
  
    it('renders main navigation links', () => {
      cy.get('nav').within(() => {
        cy.contains('Home').should('exist');
        cy.contains('Shop').should('exist');
        cy.contains('Customize').should('exist');
        cy.contains('About').should('exist');
        cy.contains('Login').should('exist');
      });
    });
  
    it('navigates to Wishlist page', () => {
      cy.contains('Wishlist').click();
      cy.url().should('include', '/wishlist');
      cy.contains('Your Wishlist').should('exist');
    });
  
    
  
    it('shows cart icon with quantity', () => {
      cy.get('nav').find('img[alt="Cart"]').should('exist');
      cy.get('nav').find('span').should('contain', '0'); // Assuming cart is empty initially
    });
  });
  