# E-Commerce Backend System

## Overview

This project implements a microservices-based backend system for a simple e-commerce application using Go and MongoDB. The system handles user authentication, product management, and order processing, emphasizing concurrency control and high availability.

## Services

### 1. User Authentication Service
- Handles user registration and login.
- Manages JWT tokens for authentication.

### 2. Product Management Service
- Handles product creation, reading, updating, and deletion.
- Manages concurrent access to product data using optimistic locking.

### 3. Order Processing Service
- Handles order creation and retrieval.
- Ensures users can only access their own orders.

## Architecture

- Microservices: Separate services for user authentication, product management, and order processing.
- Database: MongoDB for storing user information, product data, and order history.
- Authentication: JWT for secure access to protected endpoints.
- Containerization: Docker for containerizing the services.
- Clustering: Kubernetes for high availability and scalability.


## Setup and Running

### Prerequisites

- Docker
- Docker Compose
- Kubernetes
- Go
- MongoDB

### Running with Docker Compose

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/ecommerce-backend.git
   cd ecommerce-backend
2. Build and run the services:
    docker-compose up --build
3. The services will be available on the following ports:

      User Authentication Service: http://localhost:8001
      Product Management Service: http://localhost:8002
      Order Processing Service: http://localhost:8003
