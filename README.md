# Microservice Sandbox

This is a simple project that focuses on microservices architecture, consisting of two main services: the Auth Service and the Product Service.

Both services are integrated using an Nginx Proxy Pass, which efficiently manages the routing between them, ensuring seamless communication and interaction.

![system design](https://raw.githubusercontent.com/alqinsidev/microservice-sandbox/dev/docs/images/systemdesign.png)

## Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Documentation](#documentation)
- [Contact](#contact)

## Features

- **Authentication Service**: The service contain 3 end point which is Login, Register, and Validate token.

- **Product Service**: The service contains two endpoints. The first endpoint returns a list of all products, while the second endpoint specifically returns only five most expensive products and five most affordable products.

## Technologies Used

- **Auth Service**: The service are written in golang with gin for http framework, and gorm for ORM.

- **Product Service**: The service are use NodeJS with NestJS framework.

## Getting Started

The project can run either natively on your local machine or deployed as docker container.


### Run Natively

To get started with this project running on your local machine, follow the steps below:

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


### Run as Docker Container

The project are ready to use as Docker container stack, to run the project as docker container you can follow the step below:

1. **Pull the docker image**: You can pull all required images using this command:

```
docker-compose pull
```

2. **Running the container stack**: After you pull all required images you can execute this command to run the stack:

```
docker-compose up -d
```
Notes: Make sure you have all environment variables setted on `.env` file on the root of project.

## Documentation
Te project is using OpenApi for API Specs with swagger ui.

For detailed information on how to use and other details, please refer to this [project documentation](http://18.138.252.135:8012/swagger).



## Contact

If you have any questions about the project, feel free to reach me at [padlanalqinsi@gmail.com].

**SALAM WARGI JABAR**
