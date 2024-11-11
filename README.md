# Ecommerce API

This is a REST API for an ecommerce platform built using Go and Gin framework, with a Postgres database backend using GORM ORM. The API provides various endpoints for managing categories, subcategories, products, orders, and variants.

## Table of Contents

1. [API Endpoints](#api-endpoints)
   1. [Category Management](#category-management)
   2. [Subcategory Management](#subcategory-management)
   3. [Product Management](#product-management)
   4. [Order Management](#order-management)
   5. [Variant Management](#variant-management)
2. [Setup](#setup)
3. [Running the API](#running-the-api)
4. Testing with Postman


---

## API Endpoints

### Category Management

- **Create Category**
  - **Endpoint**: `POST /categories`
  - **Description**: Creates a new product category.
  - **Request Body**:
    ```json
    {
      "category_name": "Electronics",
      "description": "Category for all electronics"
    }
    ```
  - **Response**: Returns a success message or error.

- **Get All Categories**
  - **Endpoint**: `GET /categories`
  - **Description**: Retrieves all categories along with their subcategories.
  - **Response**:
    ```json
    [
      {
        "category_name": "Electronics",
        "subcategories": [
          {"name": "Smartphones"},
          {"name": "Laptops"}
        ]
      }
    ]
    ```

- **Get Category by ID**
  - **Endpoint**: `GET /categories/:id`
  - **Description**: Retrieves a specific category by its ID, along with its subcategories.
  - **Response**:
    ```json
    {
      "category_name": "Electronics",
      "subcategories": [
        {"name": "Smartphones"},
        {"name": "Laptops"}
      ]
    }
    ```

- **Update Category**
  - **Endpoint**: `PUT /categories/:id`
  - **Description**: Updates an existing category by its ID.
  - **Request Body**:
    ```json
    {
      "category_name": "Electronics & Gadgets",
      "description": "Updated description"
    }
    ```
  - **Response**: Returns the updated category.

- **Delete Category**
  - **Endpoint**: `DELETE /categories/:id`
  - **Description**: Deletes a category by its ID, and also deletes any associated subcategories.
  - **Response**: Returns a success message.

### Subcategory Management

- **Create Subcategory**
  - **Endpoint**: `POST /subcategories`
  - **Description**: Creates one or more subcategories under a parent category.
  - **Request Body**:
    ```json
    {
      "parent_category_name": "Electronics",
      "sub_category_name": ["Smartphones", "Laptops"]
    }
    ```
  - **Response**: Returns a success message.

### Product Management

- **Create Product**
  - **Endpoint**: `POST /products`
  - **Description**: Creates a new product.
  - **Request Body**:
    ```json
    {
      "product_name": "iPhone 13",
      "description": "Latest Apple iPhone",
      "price": 999.99,
      "product_category_name": "Electronics",
      "stock": 50
    }
    ```
  - **Response**: Returns a success message.

- **Get All Products**
  - **Endpoint**: `GET /products`
  - **Description**: Retrieves all products.
  - **Response**:
    ```json
    [
      {
        "product_name": "iPhone 13",
        "price": 999.99,
        "stock": 50,
        "category": "Electronics"
      }
    ]
    ```

- **Get Product by ID**
  - **Endpoint**: `GET /products/:id`
  - **Description**: Retrieves a specific product by its ID.
  - **Response**:
    ```json
    {
      "product_name": "iPhone 13",
      "price": 999.99,
      "stock": 50,
      "category": "Electronics"
    }
    ```

- **Update Product**
  - **Endpoint**: `PUT /products/:id`
  - **Description**: Updates the details of an existing product.
  - **Request Body**:
    ```json
    {
      "new_product_name": "iPhone 13 Pro",
      "description": "Latest Apple iPhone Pro",
      "price": 1099.99,
      "stock": 30
    }
    ```
  - **Response**: Returns the updated product.

- **Delete Product**
  - **Endpoint**: `DELETE /products/:id`
  - **Description**: Deletes a product by its ID.
  - **Response**: Returns a success message.

### Order Management

- **Create Order**
  - **Endpoint**: `POST /orders`
  - **Description**: Creates a new order with the selected items.
  - **Request Body**:
    ```json
    {
      "user_id": 1,
      "shipping_address": "123 Street Name, City",
      "items": [
        {"product_id": 101, "variant_id": 1, "quantity": 2, "price": 999.99}
      ]
    }
    ```
  - **Response**: Returns a success message with order ID.

- **Get All Orders**
  - **Endpoint**: `GET /orders`
  - **Description**: Retrieves all orders placed on the platform.
  - **Response**:
    ```json
    [
      {
        "order_id": 12345,
        "total_amount": 1999.98,
        "shipping_address": "123 Street Name, City"
      }
    ]
    ```

### Variant Management

- **Create Variant**
  - **Endpoint**: `POST /variants`
  - **Description**: Creates a new variant for a product.
  - **Request Body**:
    ```json
    {
      "product_name": "iPhone 13",
      "color": "Black",
      "size": 128,
      "price": 999.99,
      "stock": 50
    }
    ```
  - **Response**: Returns a success message.

- **Get All Variants**
  - **Endpoint**: `GET /variants`
  - **Description**: Retrieves all product variants.
  - **Response**:
    ```json
    [
      {
        "variant_id": 1,
        "product_name": "iPhone 13",
        "color": "Black",
        "size": 128,
        "price": 999.99,
        "stock": 50
      }
    ]
    ```

---

## Setup

### Requirements

- Go 
- Postgres
- GORM (for ORM operations)
- Gin (web framework)

### Clone Repository

git clone https://github.com/Haritha200111/E-CommerceApi/tree/development
cd repository


### Installing Dependencies

To install the necessary dependencies, run:
------go mod tidy

### Configuration

In the Go code, update the database connection details (URL, path, username, and password) as needed.

## DB Setup

Create tables in PostgreSQL using the DB scripts in the repo.

##  Running the API
###  Starting the Server
###  To start the server, run:
------go run main.go




## Testing with Postman
### Use the Postman collection provided in the repository to test the API endpoints:

### Import the Postman Collection:
Go to File > Import in Postman.
Select the E-Commerce APIs.postman_collection.json file located in the repository.

### Configure Postman Environment for placing orders:
After running the Login API, copy the authorization token and set it in the Authorization header for placing orders.

### Run API Requests:
You can now use the imported collection to test all available endpoints.
