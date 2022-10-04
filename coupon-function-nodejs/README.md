# Project description

This is an example NodeJS application, which uses Express to expose a single route.

# How to build and run the project

1. Build a Docker image

```
docker build . -t <image-name>:<version>
```

2. Push the image to a given repo, e.g. on Docker

```
 docker push <image-name>:<version>
 ```

 3. Run the application

 ```
 docker run -p 44162:3000 -d <image-name>:<version>
 ```

 # How to test it

 ```
 curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"xyz","password":"xyz"}' \
  http://localhost:44162/coupons/issued
 ```
