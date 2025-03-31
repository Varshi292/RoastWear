/// <reference types="cypress" />

describe('Wishlist Page', () => {
    const wishlistState = {
      wishlist: JSON.stringify({ items: [1] }), // ID 1 exists in your products.js
      cart: JSON.stringify({ items: [] }),
    };
  
    beforeEach(() => {
      cy.visit('/', {
        onBeforeLoad(win) {
          win.localStorage.setItem('persist:root', JSON.stringify(wishlistState));
        },
      });
  
      cy.wait(200); // Allow store rehydration
      cy.visit('/wishlist');
      cy.wait(200); // Wait to allow component to re-render after hydration
    });
  
    it('displays the wishlist title', () => {
      cy.contains('Your Wishlist ❤️').should('exist');
    });
  
    it('shows the wishlisted product', () => {
      cy.contains('I Choose Violence Funny Duck T-Shirt', { timeout: 6000 }).should('exist');
      cy.get('button').contains('Add To Cart').should('exist');
    });
  
    it('navigates to product detail page from wishlist', () => {
      cy.contains('I Choose Violence Funny Duck T-Shirt').click();
      cy.url().should('include', '/tobe-fonseca');
      cy.contains('PRODUCT DETAIL').should('exist');
    });
  
    it('shows empty message if wishlist is empty', () => {
      const emptyWishlist = {
        wishlist: JSON.stringify({ items: [] }),
        cart: JSON.stringify({ items: [] }),
      };
  
      cy.visit('/', {
        onBeforeLoad(win) {
          win.localStorage.setItem('persist:root', JSON.stringify(emptyWishlist));
        },
      });
  
      cy.visit('/wishlist');
      cy.contains('Your wishlist is empty.').should('exist');
    });
  });
  