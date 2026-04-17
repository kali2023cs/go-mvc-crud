# 🚀 Go MVC CRUD Project: The Ultimate Developer Guide

Welcome to the **Go MVC CRUD** project! This repository contains a clean, production-ready RESTful API built with Go. It features a robust **Layered Architecture**, **DTO Separation**, and **Centralized Error Handling**.

---

## 🏗️ Integrated Technologies & Architecture

-   **Gin Gonic**: High-performance HTTP web framework for routing and middleware.
-   **GORM**: Powerful Object-Relational Mapper (ORM) for Go, used for interacting with PostgreSQL.
-   **PostgreSQL**: A robust, open-source relational database.
-   **Layered Architecture**: A professional design pattern that separates concerns:
    -   **Handlers (Controllers)**: Pure HTTP logic—request parsing and response sending (`controllers/`).
    -   **Services**: Core business logic and DTO-Model mapping (`services/`).
    -   **Repositories**: Encapsulated database operations using GORM (`repositories/`).
-   **DTO (Data Transfer Objects)**: Separate structs for API requests and responses (`dto/`), ensuring internal database models remain isolated from the public contract.
-   **Centralized Error Handling**: A unified system (`utils/errors.go`) to manage and return standardized API errors.
-   **Environment Configuration (`.env`)**: Securely managing database credentials via `godotenv`.
-   **Auto Migrations**: Automatically synchronizing Go structs with database tables.

---

## 🔄 The Full Flow: Step-by-Step Logic

### 1. Fetch All Products (`GET /api/products`)
-   **Logic**:
    1.  **Handler**: Receives request, extracts query params (`name`, `min_price`, `max_price`).
    2.  **Service**: Calls repository to fetch data and maps models to `ProductResponse` DTOs.
    3.  **Repository**: Executes filtered GORM query.
    4.  **Response**: Returns `200 OK` with a list of formatted products.

### 2. Create a Product (`POST /api/products`)
-   **Input**: `name`, `description`, `price` (as JSON).
-   **Logic**:
    1.  **Handler**: Binds JSON to `CreateProductRequest` DTO and validates.
    2.  **Service**: Converts DTO to `Product` model and saves via repository.
    3.  **Repository**: Performs `Create` operation.
    4.  **Response**: Returns `201 Created` with the new product DTO.

### 3. Get Single Product (`GET /api/products/:id`)
-   **Logic**:
    1.  **Handler**: Parses ID from URL.
    2.  **Service**: Fetches product from repository.
    3.  **Error**: If not found, `utils.HandleError` returns a standardized `404 Not Found`.
    4.  **Response**: Returns `200 OK` with the product DTO.

### 4. Update Product (`PUT /api/products/:id`)
-   **Input**: JSON payload with updated fields (`UpdateProductRequest` DTO).
-   **Logic**:
    1.  **Handler**: Binds and validates incoming fields.
    2.  **Service**: Orchestrates the update through the repository.
    3.  **Response**: Returns `200 OK` with the updated product.

### 5. Delete Product (`DELETE /api/products/:id`)
-   **Logic**:
    1.  **Handler**: Triggers deletion through service.
    2.  **Service**: Ensures existence and calls repository delete.
    3.  **Response**: Returns `204 No Content`.

---

## 🎓 Go Backend Interview Questions (Updated)

### 🟢 Level 1: Architecture
1.  **Why use a Layered Architecture (Handler -> Service -> Repository)?**
    *   *Answer:* It promotes **Separation of Concerns**. Handlers only care about HTTP, Services handle logic, and Repositories handle data. This makes the code easier to test, maintain, and swap components (e.g., changing databases).
2.  **What is a DTO and why is it important?**
    *   *Answer:* Data Transfer Object. It separates the API's public contract from the internal database schema. This allows you to rename DB columns or hide sensitive fields (like `DeletedAt`) without breaking the API.

### 🟡 Level 2: Error Handling & Validation
3.  **What is the benefit of Centralized Error Handling?**
    *   *Answer:* It ensures every error response follows the same structure, making it easier for frontend developers to consume. It also reduces boilerplate code in controllers.
4.  **Explain the difference between `ShouldBindJSON` and manual parsing.**
    *   *Answer:* `ShouldBindJSON` automatically unmarshals JSON into a struct AND applies validation rules (tags) in a single step.

### 🔴 Level 3: Advanced GORM
5.  **How do you handle "Soft Deletes" in this project?**
    *   *Answer:* The `Product` model includes `gorm.DeletedAt`. When `Delete` is called, GORM sets a timestamp instead of removing the row.
6.  **How does Dependency Injection work in this project?**
    *   *Answer:* In `routes/api.go`, we initialize the Repository, then inject it into the Service, which is then injected into the Handler. This avoids global state and improves testability.

---

## 🛠️ Project Checklist
- [x] Layered Architecture (Handler/Service/Repo).
- [x] DTO for Request/Response separation.
- [x] Centralized Error Handling system.
- [x] Proper HTTP Status Codes (201, 204).
- [x] Unit testing for layered logic.
- [x] Custom validation logic.
