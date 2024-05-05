## About The Project



This project was made for improving understanding in Back-End part using Gin as a main web-framework.
<br>


### Built With
* [Golang]
* [Gin Gonic Framework]
* [Gorm]
* [Redis]
* [Docker]
* [PostgreSQL]
* [JWT Authorization]


### Layout
```tree
├── app
│   ├── controller
│   │   ├── category_controller.go
│   │   ├── product_controller.go
│   │   ├── cart_controller.go
│   │   ├── NewController.go
│   │   ├── product_controller.go
│   │   └── user_controller.go
│   │
│   ├── middleware
│   │   ├── auth.go
│   │   ├── LoadLocalVariables.go
│   │   ├── basic_auth.go
│   │   └── logger.go
│   │
│   └── server.go
│
├── src
│   ├── app
│   │   └── service
│   │       ├── category_service.go
│   │       ├── product_service.go
│   │       ├── NewService.go
│   │       ├── cart_service.go
│   │       └── user_service.go
│   │
│   ├── domain
│   │   └── model
│   │       ├── cart.go
│   │       ├── product.go
│   │       ├── category.go
│   │       └── user.go
│   │
│   └── infrastructure
│       ├── cache
│       │   ├── product-cache.go
│       │   └── redisCache.go
│       │
│       └── repository
│           ├── category_repository.go
│           ├── user_repository.go
│           ├── cart_repository.go
│           ├── product_repository.go
│           └── postgres.go
│
├── .gitignore
├── .env
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── GolangProject_postman.json
├── wait-for-postgres.sh
└── README.md
```


