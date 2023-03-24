# Test Project

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

## Table of Content

  - [About](#about)
    - [Features](#features)
    - [Design Pattern](#design-pattern)
    - [API Documentation](#api-documentation)
    - [Unit Test](#unit-test)
  - [Getting started](#getting-started)
    - [Installing](#installing)
    - [Layout](#layout)

## About

This project is example of RESTful API to store and get products.

### Features

- [x] Create/Save Product
- [x] Get Products Sorted

### Design Pattern

![mc design architecture drawio](https://user-images.githubusercontent.com/38873729/227397204-938fd524-e026-40da-b8bb-ef9fe6707f07.png)

MVC (Model-View-Controller) is one of the most commonly used software architectural designs in web and desktop application development. This design separates the application into three main components, namely the model, view, and controller. However, in this project development, the view aspect is not involved, so the architecture design becomes model MC.

Model: The model component in this design is a part of the application that handles business logic and interacts with the database. The model also serves as a representation of data in the application.

Controller: The controller component is responsible for controlling the interaction between the model and view. The controller receives requests from users and calls the model to retrieve data, then responds to the user by displaying the requested information in the view.

The advantage of using this design is separating the presentation, business logic, and user interaction, making it easier to develop, maintain, and test the application. In addition, it allows team development and work separation. In the model section, the model is responsible for handling product information. In the controller section, the controller is responsible for controlling user interaction and calling the model. This design applies the principle of "Fat model, thin controller". This means that the model should have complex business logic, and the controller only acts as a controller between the model and view. This ensures that the model has clear responsibilities and is easy to maintain. 

This design allows developers to change parts of the application without affecting other parts. For example, developers can change the presentation without modifying the business logic in the model or controller. By separating tasks into two parts, testing can be done separately for each part. For example, testing can be done to ensure that the model works correctly without having to think about the controller. Then, the application can be easily scaled. For example, if the application becomes more complex, the controller can be upgraded to handle more requests without modifying the model. 

### API Documentation

My Application Programming Interface Documentation is available at [OPEN API page](https://app.swaggerhub.com/apis/T6549/TestProducts/1.0.0) or [API page.](API.md)

### Unit Test

Coverage total unit test of this project is 91%.

<img width="703" alt="Screenshot 2023-03-23 at 21 56 56" src="https://user-images.githubusercontent.com/38873729/227323969-4ab7d042-357b-4695-b623-aa6b459667cc.png">

## Getting Started

Below I describe how to start this project

### Installing

You must download and install `Go`, follow [this instruction](https://golang.org/doc/install) to install.

After Golang installed, Follow this instructions
```bash
$ git clone https://github.com/ryanpriatama/testBackendProduct
$ go run main.go
```

Go to `http://localhost:3000/` to [start this application.](http://localhost:3000/)

### Layout

```tree
├── app
│   ├── database.go
│   ├── redis.go
│   ├── database_test.go
│   ├── redis_test.go
├── controller
│   ├── product_controller_impl.go
│   ├── product_controller.go
│   ├── product_controller_impl_test.go
├── exception
│   ├── error_handler.go
│   ├── error_handler_test.go
├── helper
│   ├── error.go
│   ├── json.go
│   ├── model.go
│   ├── tx.go
│   ├── type_const.go
│   ├── error_test.go
│   ├── json_test.go
│   ├── model_test.go
│   ├── tx_test.go
├── middleware
│   ├── auth_middleware.go
│   ├── auth_middleware_test.go
├── mocks
│   ├── mocks.go
│   ├── mocks_test.go
├── model
│   ├── domain
│   │   └── product.go
│   ├── web
│   │   └── product_create_request.go
│   │   └── product_response.go
│   │   └── web_response.go
├── repository
│   │── product_repository.go
│   │── product_repository_impl.go
│   │── product_repository_impl_test.go
├── service
│   │── product_service.go
│   │── product_service_impl_test.go
│   │── product_service_impl.go
│   │── product_service_impl_test.go
├── test
│   │── testdb.go
│   │── testdb_test.go
├── cover.out
├── API.md
├── apispec.json
├── go.mod
├── go.sum
├── main.go
├── main_test.go
├── README.md
```
