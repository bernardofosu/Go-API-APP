# 📁 my-go-gin-app/ — Gin Version of Your Express Structure

```bash
my-go-gin-app/
│
├── 📁 config/               # Configuration (env loading, DB setup)
│   ├── config.go           # Loads .env and app config
│
├── 📁 controllers/         # Route handlers (similar to Express controllers)
│   ├── userController.go
│   ├── movieController.go
│
├── 📁 models/              # Structs for database models
│   ├── user.go
│   ├── movie.go
│
├── 📁 routes/              # Route grouping and handler binding
│   ├── userRoutes.go
│   ├── movieRoutes.go
│   └── router.go           # Init all routes
│
├── 📁 utils/               # Helpers and error handling
│   ├── apiFeatures.go
│   ├── customError.go
│
├── 📁 middleware/          # Custom middleware (auth, logging, etc.)
│   ├── authMiddleware.go
│   ├── errorHandler.go
│
├── 📁 .env                 # Env variables (e.g., PORT=8080)
├── 📄 .gitignore
├── 📄 go.mod               # Go module
├── 📄 go.sum
├── 📄 main.go              # Entry point
```

---

## 🧩 Example Code for Key Files

### main.go
```go
package main

import (
    "log"
    "my-go-gin-app/config"
    "my-go-gin-app/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    r := gin.Default()

    routes.RegisterRoutes(r)

    port := config.AppConfig.Port
    log.Printf("Server running on port %s", port)
    r.Run(":" + port)
}
```

---

### config/config.go
```go
package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port string
}

var AppConfig Config

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    AppConfig = Config{
        Port: os.Getenv("PORT"),
    }
}
```
.env File: This is a file that you typically keep in the root of your project (e.g., .env). It contains key-value pairs

.env in config/ folder:
```go
err := godotenv.Load("config/.env")
```
1. .env in a parent directory:
```go
err := godotenv.Load("../.env")
```

You can even load multiple .env files:

```go
err := godotenv.Load("config/.env", "secrets.env")
```

### .env
```ini
PORT=8080
```

---

### routes/router.go
```go
package routes

import (
    "github.com/gin-gonic/gin"
    "my-go-gin-app/routes/userRoutes"
    "my-go-gin-app/routes/movieRoutes"
)

func RegisterRoutes(r *gin.Engine) {
    userRoutes.Register(r)
    movieRoutes.Register(r)
}
```

---

### routes/userRoutes.go
```go
package userRoutes

import (
    "github.com/gin-gonic/gin"
    "my-go-gin-app/controllers"
)

func Register(r *gin.Engine) {
    userGroup := r.Group("/api/users")
    {
        userGroup.GET("/", controllers.GetUsers)
    }
}
```

---

### controllers/userController.go
```go
package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    users := []string{"Alice", "Bob"}
    c.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data":   users,
    })
}
```

---

## ✅ Run It

### Install dependencies:
```bash
go mod init my-go-gin-app
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
```

### Then:
```bash
go run main.go
```

Your app will be running on:
👉 `http://localhost:8080`

Example API URL:
✅ `http://localhost:8080/api/users`

You can test it with:
```bash
curl http://localhost:8080/api/users
```

```sh
my-go-gin-app/
├── routes/                    # Main folder for all route-related files
│   ├── router.go              # Registering and initializing routes
│   └── userRoutes/            # Folder for user-specific routes
│       └── userRoutes.go      # Define user-related routes here
├── controllers/               # Controller logic (like handling requests)
│   ├── userController.go
├── config/                    # Configuration (loading .env, app settings)
├── .env                       # Environment variables (like PORT)
├── main.go                    # Entry point for the application
├── go.mod                     # Go module definition
└── go.sum                     # Go sum for dependencies
```

```sh
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get github.com/joho/godotenv
```