# Go frameworks

## http

## chi

[Chi site](https://go-chi.io)

```sh
go mod init github.com/mariolazzari/go-frameworks
go get -u github.com/go-chi/chi/v5
```

### Pros

- Light
- Fast
- No external deps
- Middlewares
- net/http compatible

### Use cases

- Short learning curve
- 

## FastHTTP

[FastHTTP site](https://github.com/valyala/fasthttp)

```sh
go get -u github.com/valyala/fasthttp
```
### Pros

- Can reuse request object
- Optimizations

### Use cases

- Video & chat servers
- Low memory usage
- Not requiring heavy business layer

## Fiber 

[Fiber site](https://github.com/gofiber/fiber)

```sh
go mod init github.com/your/repo
go get -u github.com/gofiber/fiber/v3
```

### Pros

- ExpressJS inspired
- Built on top of FastHTTP
- API servers

### Cons

- Too many deps
- Not stdlib based

### Use cases

- Streaming
