# Go Microservice Code Review Guidelines

These guidelines are meant to ensure the quality, maintainability, and consistency of our Go microservice. All code submitted for review must adhere to these principles.

## 1. Single Responsibility Principle (SRP)

Each function, method, or struct should have one, and only one, reason to change.

* **Functions and Methods:**
    * Avoid monolithic functions that handle multiple concerns (e.g., fetching data, processing it, and handling the API response).
    * A function should do one thing and do it well.

* **Structs:**
    * A struct should represent a single concept. For example, a `User` struct should not contain fields related to the HTTP request itself.

* **Packages:**
    * Each package should have a single, well-defined purpose. For example, the `repository` package should only be concerned with data persistence, not business logic.