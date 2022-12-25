# Hacktiv8-Final-Project-3

## Overview
This project is deployed at [URL](https://hacktiv8-final-project-3-production-c895.up.railway.app/). Check Out Our API [Docs](https://hacktiv8-final-project-3-production-c895.up.railway.app/docs/index.html#/).

## Getting Started
1. Clone this project or run ```git clone https://github.com/Zaidannzzz/Hacktiv8-Final-Project-3.git```
2. `cd Hacktiv8-Final-Project-3`
3. if you have **npm** installed just run `./Makefile` or `go run main.go`

## Project Structure
This project using Clean Architecture which are define as follows:

```
config/gorm.go // Database Configuration
|-- httpserver
|   |-- controllers
|   |   |-- user_controller.go
|   |   |-- category_controller.go 
|   |-- dto
|   |   |-- user_dto.go
|   |   |-- category_dto.go 
|   |-- middleware
|   |   |-- auth_middleware.go
|   |-- models
|   |   |-- user_models.go
|   |   |-- category_models.go
|   |-- repositories
|   |   |-- user_repository.go
|   |   |-- category_repository.go
|   |-- routers
|   |   |-- user_routers.go
|   |   |-- category_routers.go
|   |-- services
|   |   |-- user_services.go
|   |   |-- category_services.go
|-- docs //Define Swagger generated by swaggo
|-- utils //Define utility like helper and response standarization
```

**Feature Driven by Contributors
1. Wahyu Setiawan
- Category
- Setup Swaggo
1. Muhammad Febri Andani
- Tasks project
1. Zaidan Zulhakim
- User
- Auth
