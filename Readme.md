# FoodBridge - Bridging gap between NGO's and Resutrant's

FoodBridge is a web platform that connects **restaurants** with **NGOs** to redistribute surplus food, reducing food waste and helping those in need. The platform enables restaurants to list surplus food, and NGOs can request available food items for collection.

## Features
- 🍽 **Food Listing**: Restaurants can list surplus food with details like quantity and expiry time.
- 🏢 **NGO Requests**: NGOs can browse available food items and request collection.
- 🔐 **Authentication**: Secure **JWT-based authentication** for restaurants and NGOs.
- 📄 **Contact Handling**: NGOs can contact restaurants for coordination.
- ⚡ **High-Quality UI**: Modern UI with animations and responsive design.
- 📊 **Analytics Dashboard**: Monitor food distribution statistics.
- 🗺 **Location-Based Matching**: Connect NGOs with the nearest restaurants.
- 📱 **Mobile-Friendly UI**: Responsive design for all devices.
- 🛠 **Admin Panel**: Manage users, food listings, and NGO requests efficiently.
- 📡 **Real-time Notifications**: NGOs and restaurants receive instant updates on new listings and requests.

---

## Tech Stack
### 🔹 Frontend:
- **HTML, CSS, JavaScript** (with animations and high-quality visuals)
- **Bootstrap & Tailwind CSS** for styling
- **Chart.js** for data visualization
- **Axios** for API requests

### 🔹 Backend:
- **Golang (Gin framework)**
- **GORM (ORM for database interactions)**
- **SQLite (Database) or PostgreSQL (Optional for production)**
- **JWT Authentication**
- **RESTful API for seamless integration**
- **WebSockets for real-time updates**

### 🔹 Additional Tools:
- **Docker** for containerization
- **CI/CD Pipelines** using GitHub Actions
- **Postman** for API testing
- **Swagger** for API documentation
- **Redis** for caching and optimizing performance

---

## Installation & Setup
### 🛠 Prerequisites
- Install **Go** (https://golang.org/dl/)
- Install **Git** (https://git-scm.com/downloads)
- Install **Docker** (https://www.docker.com/)
- Install **Node.js** (if working on frontend enhancements)

### 🚀 Steps to Run the Project
#### 1️⃣ Clone the Repository:
```bash
 git clone https://github.com/saiaryansahoo/FoodBridge-web.git
 cd FoodBridge-web
```

#### 2️⃣ Install Dependencies:
```bash
go mod tidy
```

#### 3️⃣ Set Up Environment Variables:
Create a `.env` file in the backend directory and add:
```env
PORT=8080
DATABASE_URL=sqlite3://foodbridge.db
JWT_SECRET=your_secret_key
```

#### 4️⃣ Run the Backend Server:
```bash
go run main.go
```

#### 5️⃣ Open `index.html` to view the frontend

#### 6️⃣ Run with Docker (Optional):
```bash
docker build -t foodbridge-web .
docker run -p 8080:8080 foodbridge-web
```

#### 7️⃣ Run Frontend Development Server (if using a framework like React or Vue):
```bash
cd frontend
npm install
npm start
```

---

## API Endpoints
| Endpoint          | Method | Description |
|------------------|--------|-------------|
| `/auth/signup`   | POST   | User signup |
| `/auth/login`    | POST   | User login  |
| `/food/list`     | GET    | Get all food listings |
| `/food/add`      | POST   | Add new food item |
| `/ngo/request`   | POST   | NGO requests food |
| `/analytics`     | GET    | View food distribution statistics |
| `/notifications` | GET    | Get real-time notifications |
| `/admin/users`   | GET    | Fetch all users (Admin only) |

---

## Testing APIs in Postman 🚀
### 1️⃣ Import API Endpoints:
- Open **Postman**.
- Click **Import** and select the provided API collection (if available) or manually add requests.

### 2️⃣ Testing Authentication:
#### Signup:
- Method: **POST**
- URL: `http://localhost:8080/auth/signup`
- Body (JSON):
```json
{
  "username": "testuser",
  "password": "password123"
}
```
- Click **Send**.
- You should receive a response confirming signup.

#### Login:
- Method: **POST**
- URL: `http://localhost:8080/auth/login`
- Body (JSON):
```json
{
  "username": "testuser",
  "password": "password123"
}
```
- Click **Send**.
- The response should contain a **JWT token**.
- Copy the token for authorization in further API requests.

### 3️⃣ Testing Food Listing API:
#### Add Food Listing:
- Method: **POST**
- URL: `http://localhost:8080/food/add`
- Headers:
  - **Authorization**: Bearer `your_token_here`
- Body (JSON):
```json
{
  "name": "Fresh Bread",
  "quantity": 10,
  "expiry": "2025-12-31"
}
```
- Click **Send**.
- You should see a success response confirming the food listing.

#### Fetch Food Listings:
- Method: **GET**
- URL: `http://localhost:8080/food/list`
- Click **Send**.
- Response should display a list of food items available.

---

## Cross-Checking Data in SQLite 🛠
### 1️⃣ Open SQLite Database:
```bash
sqlite3 foodbridge.db
```

### 2️⃣ List Available Tables:
```sql
.tables
```
Expected Output:
```
users  food_items  ngo_requests
```

### 3️⃣ Fetch All Users:
```sql
SELECT * FROM users;
```

### 4️⃣ Fetch Food Listings:
```sql
SELECT * FROM food_items;
```

### 5️⃣ Check NGO Requests:
```sql
SELECT * FROM contacts;
```
---

## Contribution Guide 🚀
1. **Fork the repository**
2. **Create a new branch**: `git checkout -b feature-branch`
3. **Commit your changes**: `git commit -m "Added new feature"`
4. **Push to the branch**: `git push origin feature-branch`
5. **Submit a pull request** 🎉

### 🔹 Contribution Rules
- Follow **Go best practices** and maintain clean code.
- Ensure **proper API documentation** before submitting new routes.
- Write **unit tests** for backend changes.
- Maintain **UI consistency** when adding frontend elements.

---

## License
This project is **MIT Licensed**.

---

## Contact
📧 **Sai Aryan Sahoo** – [LinkedIn](https://linkedin.com/in/saiaryansahoo)  
🔗 **GitHub**: [saiaryansahoo](https://github.com/saiaryansahoo)  


## Link to Mobile App - [FoodBridge](https://github.com/saiaryansahoo/Food-Bridge)
