# Go Web Service

This repo contains a simple web service example written in Go.

## Beginning

These instructions will help you to run and develop the project on your local machine.

### Requirements

The following software must be installed to run this project:

- Go

### Installation

1. Clone the repo:

    ```bash
    git clone https://github.com/abdullahtopall/go-web-service.git
    ```

2. Go to the project directory:

    ```bash
    cd go-web-service
    ```

3. Install dependencies:

    ```bash
    go get github.com/gin-gonic/gin v1.9.1
	go get github.com/go-playground/assert/v2 v2.2.0
	go get github.com/joho/godotenv v1.5.1
	go get github.com/stretchr/testify v1.9.0
	go get github.com/swaggo/files v1.0.1
	go get github.com/swaggo/gin-swagger v1.6.0
	go get github.com/swaggo/swag v1.16.3
	go get gorm.io/driver/postgres v1.5.6
	go get gorm.io/gorm v1.25.7
    ```

4. Start the application:

    ```bash
    go run main.go
    ```

5. Go to `http://localhost:8080` in your browser and view the app.




