
# GoLang Practice

This repository contains Go projects demonstrating various features and best practices of Go programming. Below is the structure of the repository and commands to run the project.

## Repository Structure

```bash
imran@imran ~/go_practice/first_project 
├── config
│   └── config.go
├── controllers
│   ├── controller.go
│   └── user_controller.go
├── database
│   └── db.go
├── first_project
├── go.mod
├── go.sum
├── main.go
├── models
│   └── user.go
├── README.md
├── routes
│   └── routes.go
├── services
└── utils
    └── response.go
```

## Commands

1. Initialize Go module:

```bash
go mod init
```

2. Install dependencies:

```bash
go mod tidy
```

3. Build the project:

```bash
go build
```

4. Run the project:

```bash
go run main.go
```

## API Endpoints

1. GET all items:

```
GET http://0.0.0.0:8080/api/v1/home
```

2. GET a specific item by ID:

```
GET http://0.0.0.0:8080/api/v1/home/:id
```

3. Create a new item:

```
POST http://0.0.0.0:8080/api/v1/home
```

4. Delete an item by ID:

```
DELETE http://0.0.0.0:8080/api/v1/home/:id
```

## Cloning the Project

To get started with this repository, clone the project using:

```bash
git clone https://github.com/emoncse/Golang_Practice.git
cd Golang_Practice
```

Happy coding!