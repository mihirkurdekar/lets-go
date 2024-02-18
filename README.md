# Lets GO!

A microservice created to learn go with help of tutorials from [Net Ninja](https://www.youtube.com/@NetNinja).

Tutorials link: https://youtu.be/wpnN3RIRSxs?si=mQjDdeOkhjVER19s 

## Pre requisites

### Install Go

Download and install from https://go.dev/dl/

### Docker

Download and install https://www.docker.com/products/docker-desktop/



## Steps to run

### Redis
```
docker run -d -p 6379:6379 redis:latest
```

### Build
buidling will install missing dependencies
```
go build
```


### Run app

```
go run main.go
```

