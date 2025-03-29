# **RoastWear**

## **Project Description**
RoastWear is a web application that brings humor, creativity, and customization into one exciting platform. Users can:
- **Browse and Purchase** pre-designed T-shirts featuring funny memes.
- **Create Custom Designs** using an intuitive design tool.
- **Import Trending Memes** directly from **Tenor** for a seamless experience.
- **Add Personalized Text and Images** to their creations.
- **Preview in Real-Time** before placing an order.

---

## **Project Members**
- **Brandon Grill**
- **Yasasri Sabbineni**
- **Varshitha Rayala**
- **Nasim Boussarhane**

---

## **Setup and Installation**

### **Steps**

1. **Clone the Repository:**
    ```bash
    git clone https://github.com/Varshi292/RoastWear.git
    cd RoastWear
    ```

2. **Set Up Environment Variables:**  
   Create a `.env` file in the root directory:
    ```env
    PORT=7777
    STATIC_FILES_PATH=./frontend/build
    DB_PATH=./db/users.db
    ```

3. **Install Backend Dependencies:**
    ```bash
    npm install
    ```

4. **Install Frontend Dependencies and Build:**
    ```bash
    cd frontend
    npm install
    npm run build
    ```

5. **Run the Application:**
    ```bash
    go run cmd/main.go
    ```

6. **Access the App:**
    - Open your browser and go to `http://localhost:7777`.