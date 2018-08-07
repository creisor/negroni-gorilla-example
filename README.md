# Negroni + Gorilla Mux with Multiple Subrouters

## The problem

I am attempting to have multiple-versioned api's (`/v1`, `/v2`, etc), which can be accomplished.  However, when I add multiple handlers to a Negroni instance, it checks each for a route and responds accordingly with each handler.

The upshot is that, for the handler which matches, I get the desired response, but for the handler which doesn't match, I get a 404, resulting in something awkward:

With version 1:
```
creisor@gm19793-740 ~ $ curl http://localhost:3000/v1/hello
Hello from v1
404 page not found
```

And version 2:
```
creisor@gm19793-740 ~ $ curl http://localhost:3000/v2/hello
404 page not found
Hello from v2
```

Ideally, I could somehow configure it so that if a route is found, the other handler does not 404.  But, of course, if a route is not found in either handler, it would respond with 404.
