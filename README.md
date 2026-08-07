# Permanent-Storage-Todo-App

## Overview

This is a simple RESTful API for a Todo App using Go and the `gin-gonic/gin` framework for routing. The app will manage todo items stored in a database (using GORM) and support basic CRUD operations (Create, Read, Update, Delete)  with features like categorization, prioritization, completion tracking, due dates, status filtering, title search, and bulk updates.

## Features

Our app has the following features:

- ### Implementing Ten Endpoints:
    - **GET `/todos`**: Retrieve all todos as a JSON array.
    - **GET `/todos/:id`**: Retrieve a specific todo by its ID.
    - **GET `/todos/category/:category`**: Retrieve all todos in a specific category.
    - **GET `/todos/status/:status`**: Retrieve all todos with a specific completion status.
    - **GET `/todos/search`**: Retrieve todos with titles containing a query parameter `q` (case-insensitive partial match).
    - **POST `/todos`**: Create a new todo with title, completed status, category, priority, and optional due date.
    - **PUT `/todos/:id`**: Update a todo’s title, completed status, category, priority, or due date by its ID.
    - **PUT `/todos/category/:category`**: Update the `completed` status of all todos in a specific category.
    - **DELETE `/todos/:id`**: Delete a todo by its ID.
    - **DELETE `/todos`**: Delete all todos in the database.

  
 - ### Input Handling:

    - For GET , PUT, DELETE requests:
        1. Validate the input ID, ensure it is a valid number and return (400, Invalid id) if it is invalid.  
    - For POST and PUT requests: 
        1. Check the JSON body and validate incoming requests and return (400, "invalid json") if the JSON structure or types are incompatible.
        2. Validate input to ensure:
          `title` is non-empty,
          `priority` is one of "Low", "Medium", or "High",
          and `dueDate`, if provided, is a valid ISO 8601 date and not in the past (relative to the current date).
        3. Return (400, "Field cannot be empty") if there is no input for required field or the field is empty.
   -  For requests that omit an optional field:
        1. it automatically defaults to its zero-value (false, NULL).
- ### Error Handling: 
     Return appropriate HTTP status codes with clear messages to explain the error:
    - 200: Successful GET, POST, PUT, or DELETE.
    - 400: Invalid JSON or empty title or invalid ID.
    - 404: Todo not found for GET, PUT, or DELETE.
- ### Permanent Data Storage
    - Replaced the temporary in-memory slice implementation with a robust postgreSQL database integration using GORM ORM, ensuring all data persists across server restarts.
    - Integrated GORM's `AutoMigrate` feature to automatically generate and maintain the database schema directly from Go structs without requiring manual SQL table creation.
    - Assigns incremental IDs to new todos (e.g., 1, 2, 3, ...).
  
- ### Clean Code & Architecture
  This project is built following clean code principles and modular architecture to ensure high maintainability, readability and scalable growth.
  
  #### Key Architectural Highlights:
  - **Separation of Concerns**: The application logic is decoupled into **distinct layers** (model, repository, service, handler, and database configurations) to keep business logic isolated and easy to test.
  -  **Interface Based Architecture**: Lower layers define clear **Go interfaces** instead of relying on fixed implementations. This decouples the layers, make it easy to replace database implementations without touching the other layers, and allows simple mocking for unit tests.
  -  **Layer-by-Layer Validation**: Data is validated at each layer before calling the next lower layer. This ensure that every layer only receives correct and safe data, detecting errors as early as possible.
  -  **Clear Error Handling**: Errors are handled at their origin, returning simple **JSON responses** with standard HTTP status codes for easy debugging.
  
## How To Run

Follow these steps to run the code on your local device:
- Installing **Go** compiler on your device.
- Downloading the **Gin** framework for routing and the **GORM** package on your device to act as a bridge between Go code and postgreSQL.   
  To download them run the following command in your terminal:
 ```
 go mod download
 ```
- Create a postgreSQL database with the following configurations:
  - Host: localhost
  - Port: 5432
  - User: postgres
  - Password: admin
  - Database Name: todo_app

  
After downloading the previous dependencies run the following command in your terminal:
```
go run .
``` 
After this command the API server will start and run locally on
`http://localhost:8080` listening for incoming HTTP requests.

### Practical Example:
- Try to send a **POST** request to `http://localhost:8080/todos` with the JSON body:
  ```
  {"title": "Read book", "category": "study", "priority": "medium", completed": false}
  ```
- The system assigns an ID (e.g., 1) and stores the todo in the database.
- The response will be:
  ```
  {
    "id": 1,
    "title": "Read book",
    "category": study,
    "priority": medium,
    "completed": false,
    "completedAt": null,
    "dueData": null
  }
  ```