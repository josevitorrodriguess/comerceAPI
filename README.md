# Products API

## Overview

This API project provides functionalities for managing products, clients, and merchants. It includes CRUD operations for products and authentication for clients and merchants through RESTful endpoints.

## ER Diagram

![ER Diagram](diagram.png)

## Table of Contents

1. [Overview](#overview)
2. [Dependencies](#dependencies)
3. [Usage Instructions](#usage-instructions)
4. [Additional Notes](#additional-notes)


## Overview

This API project provides functionalities for managing products, clients, and merchants. It includes CRUD operations for products and authentication for clients and merchants through RESTful endpoints.

## Directory Structure

- controller: Contains HTTP request handlers for different entities.
- model: Defines data structures for login, client, and merchant.
- repository: Implements data access logic for clients and merchants.
- routes: Defines API routes and their associations with controllers.
- secase: Implements business logic operations for authentication and entity management.
- 
## Dependencies

- Gin (github.com/gin-gonic/gin) for HTTP routing.
- MySQL (github.com/go-sql-driver/mysql) for MySQL database connection.

## Usage Instructions

### Clone the repository:

```
git clone https://github.com/your-username/your-project.git
cd your-project
```
### Instal Dependencies:

```
go mod tidy
```

### Configure Database:

```
dbConnection, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
```
### run application:

```
go run main.go
```
### Additional Notes

  - The project follows an MVC (Model-View-Controller) architecture, clearly separating responsibilities among different components.
  - For security reasons, ensure that database credentials and JWT tokens are managed securely and not exposed in public repositories.


