describe('Authentication Tests', () => {
  beforeEach(() => {
    cy.visit('http://localhost:3000'); // Adjust the URL as per your setup
  });

  context('Login Page Tests', () => {
    it('should display login form correctly', () => {
      cy.visit('/login');
      cy.get('h2').contains('Login');
      cy.get('input[placeholder="Username"]').should('be.visible');
      cy.get('input[placeholder="Password"]').should('be.visible');
      cy.get('button').contains('Login').should('be.visible');
    });

    it('should show error for empty fields', () => {
      cy.visit('/login');
      cy.get('button').contains('Login').click();
      cy.contains('Login failed').should('be.visible');
    });

    
  });

  context('Register Page Tests', () => {
    it('should display register form correctly', () => {
      cy.visit('/register');
      cy.get('h2').contains('Register');
      cy.get('input[placeholder="Username"]').should('be.visible');
      cy.get('input[placeholder="Email"]').should('be.visible');
      cy.get('input[placeholder="Password"]').should('be.visible');
      cy.get('button').contains('Register').should('be.visible');
    });

    it('should show error for empty fields', () => {
      cy.visit('/register');
      cy.get('button').contains('Register').click();
      cy.contains('All fields are required!').should('be.visible');
    });

    
  });
});
