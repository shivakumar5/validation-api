# validation-api
The repo contains the code to validate the input string using RESTful Web Service

## Use Case
It will receive a string as input, potentially a mixture of upper and lower case, numbers, special characters etc. The task is to determine if the string contains at least one of each letter of the alphabet. Return true if all are found and false if not.

## Solution

Used **Gorilla Mux** package to implement a request router and dispatcher for matching incoming requests to their respective handler.

Import it with below URL:

```
github.com/gorilla/mux
```

## This API Contains the below endpoint:

```
localhost:9090/validate
```

## Validating the Input String
when user provides the input string, the below code validates it 

```go
for _, c := range key {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		default:
			return false
		}
```

when user makes a request with input string which contains **at least one special charater and alpha numeric value** , the response will be `true` else `false`.



## Swagger Document

The below swagger.json file has beenautomatically generated using `Swagger Inspector`.

```json
openapi: 3.0.1
info:
  title: ValidateAPI
  description: API to Validate the given string
  version: '1.0'
servers:
  - url: 'http://localhost:9000'
paths:
  /validate:
    get:
      description: API to Validate the given string
      parameters:
        - name: inputstring
          in: query
          schema:
            type: string
          example: Shiva!40123
      responses:
        '200':
          description: API to Validate the given string
          content:
            text/plain; charset=utf-8:
              schema:
                type: string
              examples: {}
      servers:
        - url: 'http://localhost:9000'
    servers:
      - url: 'http://localhost:9000'
```
you can find the `swagger.json` file in the main repo.

## Deploy the Service with AWS EKS and ECR

We need to follow the below steps to deploy this service to Kubernetes using AWS EKS and ECR images.

1. Create a RESTful API service using Golang
   
2. Create a Docker image of our service using Dockerfile
   
```dockerfile
FROM go_base:latest

EXPOSE 8080

WORKDIR /srv

ENTRYPOINT ["/srv/bin/validation_api"]

ENV APP_NAME=validation_api \
    AUTH_ENV=development \
    HOME=/srv

# Binaries
RUN mkdir ./bin
COPY validation_api ./bin/validation_api

# Build as root; run as unprivileged user
USER www-data
```
Run Docker build command to create image

```
docker build -t validate-api .
```

3. Create an ECR Repository
   
   Before we can push the docker image, we need to create a repository on ECR. Using AWS Console, we can create a repository. After creating an ECR repository, to push the docker image we need to authenticate our AWS CLI and tag the docker image.

   And, using below command, push image to ECR Repository

```
docker push 123456-dkr.ecr.us-east-1.amazonaws.com/one-technopedia/validation-api:latest
```

4. Create AWS EKS Cluster

    Creating an EKS Cluster is a big process, need to configure many things here. we can create EKS Cluster using Terraform also.

5. Create Deploy Manifest
   
   Once the cluster is configured and it is up and running, next step is to apply the service file to create deployment for our application/service.

   Below is the deployment and service manifest that will be used to run our application.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: validation-api
  namespace: default
spec:
  replicas: 2
  selector:
    matchLabels:
      app: validation-api
  template:
    metadata:
      labels:
        app: validation-api
    spec:
      containers:
      - name: validation-api
        image: 123456-dkr.ecr.us-east-1.amazonaws.com/one-technopedia/validation-api:latest
---
apiVersion: v1
kind: Service
metadata:
  name: validation-api
  namespace: default
spec:
  type: NodePort
  ports:
  - port: 80
    targetPort: 8080
```
* Our application label will be `app:validation-api`

* In `spec:template:spec:containers` set image for the AWS ECR image we pushed

* Number of replicas for the application is 2

Run the following command to create our deployment.

```
kubectl apply -f validation_api.yaml
```

Check the deployment and service by running below commands:

```
kubectl get deployments
kubectl get svc
```

* service type is **Nodeport**
* targetPort is 8080 since that is our container exposed port
* the selector will be app:validation-api since that is the label we defined in our deployment


6. Access the service
   
   To locally access the service, port-foward the service using below command:

```
kubectl port-forward svc/validation-api 8080:80
```

Note: If you would like to access the service from cluster, we need to allow inbound rules in security group for the nodeport IP Adress port.






















