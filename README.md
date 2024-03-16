# GO REST Backend
This is a simple starter project that helps you get started with developing RESTFUL api backends in Go

Following features are provided:
* RESTFUL endpoints
* Configuration management
* Logging with contextual information
* Error handling
* Full test coverage
* Live reloading in development
* Containerization

Following thirdparty go packages are used in the starter project
* Routing: [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [zap](https://github.com/uber-go/zap)
* JWT: [jwt-go](https://github.com/dgrijalva/jwt-go)


# Build the server
make build

# Run the server
make run

At this time, you have a RESTful API server running at `http://localhost:8080`. It provides the following endpoints:

* `GET /health`: a health check endpoint for distributed environments (Helps when running this server in kubernetes) 
* `GET /metrics`: prometheus metrics that can be scrapped by prometheus server

make run-live

You can try accessing the healthcheck endpoint from a browser or postman or terminal at http://localhost:8080/health
also try accessing the metrics endpoint from a browser or postman or terminal at http://localhost:8080/metrics

# Build container image using PODMAN
make build-container
