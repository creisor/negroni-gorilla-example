# Negroni + Gorilla Mux with Multiple Subrouters

This demonstrates how to have multiple-versioned api's (`/v1`, `/v2`, etc) using [Gorilla Mux](https://github.com/gorilla/mux) and [Negroni](https://github.com/urfave/negroni).

    curl http://localhost:3000/v1/hello

or

    curl http://localhost:3000/v2/hello


The idea of a `routes.go` is inspired by [this article](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)
