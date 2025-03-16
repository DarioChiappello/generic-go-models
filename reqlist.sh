# Users

curl -X POST http://localhost:8080/users   -H "Content-Type: application/json"   -d '{"name": "John Doe", "email": "john@example.com"}'

curl -X GET http://localhost:8080/users

curl -X GET http://localhost:8080/users/1

curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "John Updated", "email": "updated@example.com"}'

  curl -X PATCH http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"email": "newemail@example.com"}'

curl -X DELETE http://localhost:8080/users/1

# Products

curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Laptop", "price": 999.99}'

curl -X GET http://localhost:8080/products

curl -X GET http://localhost:8080/products/1

curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Laptop Pro", "price": 1499.99}'

curl -X PATCH http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{"price": 799.99}'

curl -X DELETE http://localhost:8080/products/1

# Orders

curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "product_id": 1, "quantity": 2}'

curl -X GET http://localhost:8080/orders

curl -X GET http://localhost:8080/orders/1

curl -X PUT http://localhost:8080/orders/1 \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "product_id": 1, "quantity": 5}'

curl -X PATCH http://localhost:8080/orders/1 \
  -H "Content-Type: application/json" \
  -d '{"quantity": 3}'

curl -X DELETE http://localhost:8080/orders/1

