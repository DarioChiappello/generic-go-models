# Generic Go CRUD API

A generic CRUD API implementation in Go using Gin, GORM, and PostgreSQL, following a layered architecture with method inheritance.

## Project Structure

```
├── models/ # Database models
│ ├── user.go
│ ├── product.go
│ └── order.go
├── repositories/ # Database operations
│ └── base_repository.go
├── services/ # Business logic
│ ├── base_service.go
│ ├── user_service.go
│ ├── product_service.go
│ └── order_service.go
├── controllers/ # HTTP handlers
│ ├── user_controller.go
│ ├── product_controller.go
│ └── order_controller.go
├── routers/ # Route definitions
│ ├── user_router.go
│ ├── product_router.go
│ └── order_router.go
└── main.go # Entry point

```
## Method Inheritance Flow


```
    CRUD Operations
            |
      Base Repository
            ↑
      Base Service
            ↑
    Specific Service (User/Product/Order)
            ↑
    Controller
            ↑
    HTTP Router
```

#### Request Flow for GET
```
HTTP Request → Router → Controller → Service → Repository → Database
                                                              ↓
Response ← JSON ← Controller ← Service ← Repository ← Query Results

```

### Key Components

1. **Base Repository**
   - Implements generic CRUD operations
   - Used by all services
   - Methods:
     ```go
     GetAll()
     GetByID()
     Create()
     Update()
     Patch()
     Delete()
     ```

2. **Base Service**
   - Wraps repository methods
   - Inherited by specific services
   - Example:
     ```go
     type BaseService[T any] struct {
         repo *BaseRepository[T]
     }
     ```

3. **Concrete Services**
   - Inherit from BaseService
   - Can add model-specific logic
   - Example (UserService):
     ```go
     type UserService struct {
         *BaseService[models.User]
     }
     ```

4. **Controllers**
   - Handle HTTP requests
   - Use service layer
   - Example UserController methods:
     ```go
     CreateUser()
     GetAllUsers()
     GetUserByID()
     UpdateUser()
     DeleteUser()
     ```

## API Endpoints

### Users
| Method | Endpoint       | Description                |
|--------|----------------|----------------------------|
| POST   | `/users`       | Create user                |
| GET    | `/users`       | Get all users              |
| GET    | `/users/{id}`  | Get user by ID             |
| PUT    | `/users/{id}`  | Full update user           |
| PATCH  | `/users/{id}`  | Partial update user        |
| DELETE | `/users/{id}`  | Delete user                |

### Products
| Method | Endpoint         | Description                |
|--------|------------------|----------------------------|
| POST   | `/products`      | Create product             |
| GET    | `/products`      | Get all products           |
| GET    | `/products/{id}` | Get product by ID          |
| PUT    | `/products/{id}` | Full update product        |
| PATCH  | `/products/{id}` | Partial update product     |
| DELETE | `/products/{id}` | Delete product             |

### Orders
| Method | Endpoint        | Description                |
|--------|-----------------|----------------------------|
| POST   | `/orders`       | Create order               |
| GET    | `/orders`       | Get all orders             |
| GET    | `/orders/{id}`  | Get order by ID            |
| PUT    | `/orders/{id}`  | Full update order          |
| PATCH  | `/orders/{id}`  | Partial update order       |
| DELETE | `/orders/{id}`  | Delete order               |

## Request Examples

### Users
**Create User**  
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Smith", "email": "alice@example.com"}'
```

Response

```json
{
  "ID": 1,
  "name": "Alice Smith",
  "email": "alice@example.com",
  "createdAt": "2025-03-16T12:00:00Z"
}
```

### Products
**Update Product** 

```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Premium Laptop", "price": 1499.99}'
```
Response

```json
{
  "ID": 1,
  "name": "Premium Laptop",
  "price": 1499.99,
  "updatedAt": "2025-03-16T12:05:00Z"
}
```
###  Orders
**Create Order** 


```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "product_id": 1, "quantity": 2}'
```
Response

```json
{
  "ID": 1,
  "user_id": 1,
  "product_id": 1,
  "quantity": 2,
  "createdAt": "2025-03-16T12:10:00Z"
}
```

## Setup & Run

Requirements

Go 1.18+

PostgreSQL

Docker (optional)

Environment Variables (.env)

env
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_SSLMODE=disable
```

Run with Docker

```bash
docker-compose up -d  # Starts PostgreSQL
go run main.go        # Starts API server
```

Manual Setup
```bash
# Install dependencies
go mod tidy

# Run migrations
DB_NAME=yourdb go run main.go migrate

# Start server
go run main.go
```

### License
MIT License - See LICENSE for details

This README:
1. Explains the layered architecture with method inheritance
2. Shows complete endpoint documentation
3. Provides ready-to-use request examples
4. Includes setup instructions
5. Maintains generic implementation approach
