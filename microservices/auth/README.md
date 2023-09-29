# Auth microservice

### Stack
 - Go
 - PostgreSQL
 - Redis
 - gRPC
 - Echo (web framework)


### Server (?)
This ms uses two ports, one for HTTP server, for things like register, and another one for gRPC server, for things like session validation from others ms.

For the HTTP server we're using [echo](https://echo.labstack.com/).
