# 🚀 Go MVC CRUD Project: The Ultimate Developer Guide

Welcome to the **Go MVC CRUD** project! This repository contains a clean, production-ready RESTful API built with Go. It features the **MVC (Model-View-Controller)** architecture, **GORM database management**, and **PostgreSQL integration**.

---

## 🏗️ Integrated Technologies & Concepts

-   **Gin Gonic**: High-performance HTTP web framework for routing and middleware.
-   **GORM**: Powerful Object-Relational Mapper (ORM) for Go, used for interacting with PostgreSQL.
-   **PostgreSQL**: A robust, open-source relational database.
-   **MVC Architecture**: A design pattern that separates the application into three main logical components:
    -   **Models**: Data structure and database logic (`models/product.go`).
    -   **Views**: JSON responses (In this API, the "View" is the JSON data sent back to the client).
    -   **Controllers**: Application logic and request handling (`controllers/product_controller.go`).
-   **Environment Configuration (`.env`)**: Securely managing database credentials and server ports via `godotenv`.
-   **Auto Migrations**: Automatically synchronizing Go structs with database tables.
-   **Soft Deletes**: Using GORM's `DeletedAt` to mark records as deleted without actually removing them from the database.
-   **Structured Error Handling**: A centralized system to convert complex validation errors into simple, readable maps for the frontend.
-   **Custom Validation Rules**: Extending the base validator with domain-specific rules (e.g., preventing reserved names).

---

## 🔄 The Full Flow: Step-by-Step Logic

### 1. Fetch All Products (`GET /api/products`)
-   **Logic**:
    1.  **Initialization**: A slice of `Product` models is declared.
    2.  **Database Query**: Connects to `config.DB` and builds the query.
    3.  **Filtering**: Supports optional query parameters `name` (LIKE match), `min_price`, and `max_price`.
    4.  **Response**: Returns the filtered data in a JSON object with a `200 OK` status.

### 2. Create a Product (`POST /api/products`)
-   **Input**: `name`, `description`, `price` (as JSON).
-   **Logic**:
    1.  **Request Binding**: `c.ShouldBindJSON` maps the incoming JSON to a Product struct.
    2.  **Tag Validation**: Standard tags like `required` and `gt=0` are checked first.
    3.  **Custom Validation**: The engine checks the `not-reserved` rule to ensure names like "Admin" are blocked.
    4.  **Error Formatting**: If any check fails, `utils.FormatError` kicks in to return a structured map of errors.
    5.  **Database Insert**: `config.DB.Create(&product)` adds the new record.
    6.  **Response**: Returns the created product with its generated ID and timestamps.

### 3. Get Single Product (`GET /api/products/:id`)
-   **Logic**:
    1.  **Lookup**: `config.DB.Where("id = ?", c.Param("id")).First(&product)` searches for the specific ID.
    2.  **Error Handling**: If `gorm.ErrRecordNotFound` occurs, it returns a `400 Bad Request` with "Record not found!".
    3.  **Response**: Returns the single product object.

### 4. Update Product (`PUT /api/products/:id`)
-   **Input**: JSON payload with updated fields.
-   **Logic**:
    1.  **Find First**: Locates the existing product by ID.
    2.  **Input Binding**: Binds the new data from the request body.
    3.  **Update**: `config.DB.Model(&product).Updates(input)` applies only the changed fields to the database.
    4.  **Response**: Returns the updated product record.

### 5. Delete Product (`DELETE /api/products/:id`)
-   **Logic**:
    1.  **Find First**: Ensures the product exists before attempting deletion.
    2.  **Soft Delete**: `config.DB.Delete(&product)` populates the `deleted_at` column.
    3.  **Response**: Returns a success message.

---

## 🎓 Go Backend Interview Questions & Answers

### 🟢 Level 1: Basics (Gin & MVC)
1.  **What does MVC stand for and how is it used here?**
    *   *Answer:* Model-View-Controller. Models define the data (Product), Controllers handle the logic (CRUD functions), and Routes define the entry points.
2.  **What is the benefit of using `gin.H`?**
    *   *Answer:* It is a shortcut for `map[string]interface{}`, allowing for quick and readable JSON response generation.
3.  **Why do we use `c.ShouldBindJSON` instead of manually parsing the body?**
    *   *Answer:* It handles the boilerplate of reading the request body, unmarshaling the JSON into a struct, and validation in one step.

### 🟡 Level 2: Intermediate (GORM & Database)
4.  **What is a "Soft Delete" in GORM?**
    *   *Answer:* By including `gorm.DeletedAt` in the struct, GORM will set a timestamp in that column instead of deleting the row. Queries will automatically exclude these rows unless `.Unscoped()` is used.
5.  **Explain `config.DB.Model(&product).Updates(input)`.**
    *   *Answer:* It specifies which record to update (`Model(&product)`) and then applies the changes from the `input` struct. Crucially, GORM `Updates` with a struct only updates non-zero fields.
6.  **How do you handle 404 errors when searching for a record?**
    *   *Answer:* We check the error returned by the `First()` method. If `err != nil`, we assume the record wasn't found (or another DB error occurred) and return a relevant error code.

### 🔴 Level 3: Advanced (Architecture & Design)
7.  **What is the purpose of `AutoMigrate`?**
    *   *Answer:* It automatically creates or updates database tables to match your Go structs. This ensures the schema is always in sync with the code during development.
8.  **How would you implement pagination in `GetProducts`?**
    *   *Answer:* I would use `.Limit()` and `.Offset()` methods in GORM, usually driven by `page` and `limit` query parameters from the request.
9.  **Why use a separate `config` package for the database?**
    *   *Answer:* It promotes "Separation of Concerns." The database connection logic is isolated from the application logic, making it easier to manage and share the `DB` instance across different packages.

### ⛓️ Level 4: Deep Dive (Performance & Security)
10. **Is this API vulnerable to SQL Injection?**
    *   *Answer:* No, because we use GORM's parameter binding (e.g., `Where("id = ?", c.Param("id"))`). This uses prepared statements, ensuring the input is treated as data, not code.
11. **How would you add validation to the Product Name (e.g., minimum 3 characters)?**
    *   *Answer:* I would add the `binding:"required,min=3"` tag to the Struct field and Gin's validator would automatically enforce it during `ShouldBindJSON`.
12. **Explain the difference between `PUT` and `PATCH`.**
    *   *Answer:* `PUT` traditionally replaces the entire resource, while `PATCH` performs partial updates. In this project, `PUT` is used with GORM's `Updates`, effectively behaving like a patch/partial update.
13. **Why use an in-memory SQLite database for unit testing?**
    *   *Answer:* It is extremely fast and ephemeral. It allows us to run a full test suite without needing a real PostgreSQL server or affecting our production data. Every test run starts with a clean database.
14. **How does Gin handle validation errors?**
    *   *Answer:* When `ShouldBindJSON` fails due to `binding` tags (like `required` or `gt=0`), it returns an error. We catch this error and return a `400 Bad Request` with a descriptive message to the client.
15. **What is the advantage of structured error responses (e.g., a map of field -> message)?**
    *   *Answer:* It makes it extremely easy for frontend developers to map errors directly to specific input fields in the UI, providing a much better user experience than a single error string.
16. **How do you register a custom validator in Gin?**
    *   *Answer:* You access the validator engine via `binding.Validator.Engine()`, cast it to `*validator.Validate`, and then call `RegisterValidation` with a unique tag name and a validation function.

---

## 🛠️ Project Checklist
- [x] CRUD implementation for Products.
- [x] PostgreSQL connection set up.
- [x] Auto-migration enabled.
- [x] Environment variables managed.
- [x] Implement Request Validation (using validator tags).
- [x] Add Search/Filter functionality to `GetProducts`.
- [x] Unit testing for Controllers.

