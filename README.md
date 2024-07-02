# Zocket_Rate_Limiting_Middleware

This is a simple rate-limiting-middleware written in Golang

## Introduction

This project implements a rate limiting middleware using Go and the Echo framework to restrict the number of requests per minute from specific IP addresses to various endpoints. It uses basic token bucket approach for rate limiting.

## Technologies Used

- Language: Go
- Framework: Echo (chosen for its simplicity, performance, and ease of use)
- Database: SQLite (intended for future use as a lightweight and easy-to-setup solution)

## Project Structure

The project is structured as follows:

rate_limiting_middleware/
│
├── config/
│   └── config.go
│
├── handlers/
│   ├── rate_limiter_handler.go
│   └── server_handler.go
│
├── middleware/
│   └── rate_limiter_middleware.go
│
├── routes/
│   ├── rate_limiter_route.go
│   └── server_route.go
│
├── tests/
│   └── rate_limiter_middleware_test.go
│
├── go.mod
├── go.sum
├── main.go
└── README.md

- `handlers/`: Contains the request handlers for each API endpoint.
- `routes/`: Defines the API routes and links them to corresponding handlers.
- `main.go`: Entry point of the application.
- `go.mod`: Go module file listing project dependencies.
- `config/`: Contains the configs related to the project
- `middleware/`: Contains the rate limiting middleware implementation.


## Installation

1. Clone the repository:

    https://github.com/harshsngh1/Zocket_Rate_Limiting_Middleware

2. Navigate to the project directory:

    cd rate_limiting_middleware

3. Run the following command to start the server:

    go run main.go

The server will start on localhost:8080 by default.

## Usage

### Setting Rate Limits
Use the following endpoint to set rate limits:

- Endpoint: /set-rate-limit
- Method: POST
- Payload: JSON object with endpoint, ip, and limit fields.

```curl -X POST -H "Content-Type: application/json" -d '{"endpoint": "/endpoint1", "ip": "127.0.0.1", "limit": 3}' "http://localhost:8080/set-rate-limit```

### Getting Rate Limits
Use the following endpoint to retrieve rate limits:

- Endpoint: /get-rate-limits
- Method: GET

```curl -X GET "http://localhost:8080/get-rate-limits"```

## Testing Rate Limiting
To test rate limiting, make multiple requests to an endpoint within a short time frame. You can also use the below command to test this manually.

```for i in {1..10}; do   curl -X GET http://localhost:8080/endpoint1   -H "X-Real-IP: 192.168.0.181"   -H "Content-Type: application/json"   -d '{"key": "value"}' &   sleep 1 &done```

You can also use the test file rate_limiting_middleware_test.go to test the rate limiting

## Running via Docker

To run the application using Docker, execute the following commands:

1. Build the Docker image:

docker build -t rate-limiting-app .

2. Run the Docker container:

docker run -p 8080:8080 rate-limiting-app

The application will start inside a Docker container and be accessible at http://localhost:8080.