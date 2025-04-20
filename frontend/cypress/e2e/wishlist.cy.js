const wishlistState = {
  cart: { items: [] },
  wishlist: { items: [1] },
  _persist: { version: -1, rehydrated: true },
};

describe('Wishlist Page', () => {
  beforeEach(() => {
    cy.visit('/', {
      onBeforeLoad(win) {
        win.localStorage.setItem('persist:root', JSON.stringify({
          cart: JSON.stringify(wishlistState.cart),
          wishlist: JSON.stringify(wishlistState.wishlist),
          _persist: JSON.stringify(wishlistState._persist),
        }));
      },
    });

    cy.visit('/wishlist');
  });

  it('displays the wishlist title', () => {
    cy.contains(/your wishlist/i).should('exist');
  });

  


  it('shows empty message if wishlist is empty', () => {
    const emptyState = {
      cart: { items: [] },
      wishlist: { items: [] },
      _persist: { version: -1, rehydrated: true },
    };

    cy.visit('/', {
      onBeforeLoad(win) {
        win.localStorage.setItem('persist:root', JSON.stringify({
          cart: JSON.stringify(emptyState.cart),
          wishlist: JSON.stringify(emptyState.wishlist),
          _persist: JSON.stringify(emptyState._persist),
        }));
      },
    });

    cy.visit('/wishlist');
    cy.contains(/wishlist is empty/i).should('exist');
  });
});