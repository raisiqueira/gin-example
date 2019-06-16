# GIN - Example

> For studies only 🙅‍♂️ don`t use it in production!

Basic example of Gin web framework on Docker

## Build the image

Build the image

`docker build -t raisiqueira/gin-go:1.0.1`

## Run

`docker container run -p 8080:8080 raisiqueira/gin-go:1.0.1`

## Running on Kubernetes

`kubectl apply -f deployment.yaml` and `kubectl apply -f service.yaml`.
Get the _External IP_ from the created service, and open it in a browser on the **8080** port.