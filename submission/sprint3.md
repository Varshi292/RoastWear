# Sprint 3 Report

## User Stories

1. **As a user,** I want to customize my own T-shirts so that I can create a unique design.
2. **As a user,** I want to add T-shirts to my cart so that I can purchase them later.
3. **As a user,** I want a secure login/signup system so that my data is protected.
4. **As a user,** I want to see my previous designs and meme uploads so my workflow is more efficient.
5. **As a user,** I want to be able to press the general box around a product, not just the image, and still be taken to the product page.

---

## Issues Planned to Address

- Enhancing an the cart Layout.
- Implementing WishList feature.
- Improving the UI/UX of the website.
- updating the About page
- Developing the T-shirt customization feature
- Implementing Cypress Component Testing on new features
- Implementing unit Testing on Cart functionalities
- Implementing unit Testing on Wishlist functionalities
- Implementing unit Testing on all the upgrades small functions.

- Integration of Go in backend.
- Create API documentation
- Create general documentation for backend
- Improving error handling in user authentication.
- Create backend unit tests for handlers
- Implement the T-shirt customization feature.
- Improve UI design based on feedback.
- Enhance authentication validation and error handling.
- More organized user data organization and retrieval (to set us up for shopping carts, past 
  designs, and past purchases).
- Fully remake backend with Go.
- Add validation of empty input outside of backend.
- Implement session management.

---

## Successfully Completed

1. Upgraded the usability of Cart function.
2. Implemented the wishlist feature.
3. Improved the UI/UX of the website.
4. Implemented Cypress Component Testing on new features.
5. Implemented unit Testing on Cart functionalities.
6. Implemented unit Testing on Wishlist functionalities.
7. Implemented unit Testing on all the upgrades small functions.
8. Improved error handling in user authentication.
9. Updating the About page.

10. Retrieval of previous meme uploads
11. Partially implemented session management.
12. Partially implemented server-side user authentication
13. Implemented bcrypt password hashing and verification
14. Created API documentation using Swagger
15. Designed unit tests for the API using Postman
16. Fully* recreated backend functionality in Go
17. Partially implemented image file storage
18. Implemented environment variables and further customizability of application
19. Improvement of User model through additional layers of abstraction

---

## Issues Not Completed and Reasons

- **User authentication**  was partially implemented due to time constraints alongside full Go integration.
- **The T-shirt customization feature** was not completed due to prioritization of authentication and UI setup.
- **Image uploading** was partially completed, but could not be completed in time due to wide specification.
- **Tokenized sessions** was put on hold due to the need to remake certain backend processes.

---

## Next Steps

- Implement the T-shirt customization feature.
- Improve UI design based on feedback.
- Enhance authentication validation and error handling.
- More organized user data organization and retrieval (to set us up for shopping carts, past designs, and past purchases).
- Add validation of empty input outside of backend.
- Implement session management.
- Combine cart with backend

  ## List of unit tests
  ## Frontend
  - Register.test.js
  - Home.test.js
  - Navbar.test.js
  - CartItem.test.js
  - CartTab.test.js
  - login.cy.js
  - navbar.cy.js
  - ProductDetail.cy.js
  - wishlist.cy.js

  ## Backend
  - Register New User
  - Register Existing Username
  - Register Existing Email
  - Register Existing Password
  - Register Empty
  - Register No Username
  - Register No Email
  - Register No Password
  - Register invalid Request
  - Login User
  - Login Wrong Username
  - Login Wrong Password
  - Login No Username
  - Login No Password
  - Login Empty
  - Login Invalid Request
 
  ## Backend for Media Management
  -Post Proper Image
  -Post Improper Image
  -Recieve all of the images a past user uploaded
  -Recieve an error for trying to recieve images from a user that doesn't exist or has no media

   ## Backend for Session Management (Steps for each test)
  - Login -> Generate Session ID -> Verify ID
  - Login -> Generate Session ID -> Verify ID -> Delete Session -> Confirm Session Deleted
  - Login -> Generate Session ID -> Verify ID -> Modify Session Token -> Verify Session Token Doesn't Work
  

  
## Backend Video

https://youtu.be/PFCtahJQkfg

## Backend API Documentation

https://drive.google.com/file/d/1YmjcxXG4b_tW9HGEoV3FlPjdADqeLdOz/view?usp=sharing

## Frontend Video

https://1drv.ms/v/c/1a653b0f309d6b9d/EdNC0FtABj1CgkAlfPIIjnoBdbJi-NLP-45QWegxyawfqQ
