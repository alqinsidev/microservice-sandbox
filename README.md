# Microservice Sandbox

This is a simple project that combines authentication (auth) and product functionality in the backend. It provides a robust foundation for managing user authentication and product data for your application.

![system design](https://github.com/alqinsidev/microservice-sandbox/blob/dev/docs/images/systemdesign.png)

## Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Documentation](#documentation)
- [Contact](#contact)

## Features

- **Authentication Service**: The service contain 3 end point which is Login, Register, and Validate token.

- **Product Service**: The service contain 2 end point which return list that sorted by price.

## Technologies Used

- **Auth Service**: The service are use golang with gin for routing, and gorm for ORM.

- **Auth Service**: The service are use golang with gin for routing, and gorm for ORM.

## Getting Started

To get started with this project, follow the steps below:

1. **Clone the repository**: Begin by cloning this repository to your local machine using the following command:

   ```
   git clone https://github.com/alqinsidev/microservice-sandbox.git
   ```

2. **Installation**: The project has 2 service, to install dependencies for the product service

   ```
   cd project-service
   npm install
   ```

3. **Configuration**: Change every `.env.example` into `.env` and adjust the value based on your environment

4. **Run the Application**: Start the application by executing the following command:


- For the auth-service
   ```
   cd auth-service
   go mod run main.go
   ```

- For the product-service
   ```
   cd product-service
   npm run start:dev
   ```

## Documentation

For detailed information on how to use and customize this project, refer to the [project documentation](http://18.138.252.135:8012/swagger).



## Contact

If you have any questions or need assistance, please feel free to contact the project maintainer at [padlanalqinsi@gmail.com].
