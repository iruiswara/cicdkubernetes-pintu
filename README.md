# CI / CD Kubernetes

this repository contain technical test for pintu
    
---

## Tasks
1. Create a simple app using both programming languages (optional)
note: you can use the boiler plate in any repository

Answers: 
+ already create 2 backend API `Hello,World` with go lang and node js
    + [backend API golang](./services/backend-go/README.md)
    + [backend API node js](./services/backend-node/README.md)

2. Create CI/CD start from:

    a. Build and test (optional) the services

    b. Building the container image
    
    c. Store container image to a registry

    d. Deploy to Kubernetes cluster

    note: this is not a mandatory step, you can provide the manifest only.

Answer:
+ i use jenkins for CI/CD and deployment along with explaination regarding the stages
    + [Deployment CI/CD](./Deployment/README.md)

3. Service can be consumed from the public internet using the DNS domain (optional)
note: you can use xip.io
+ you can try access backend API from broswer in this URL: [https://8714-108-137-148-243.ngrok-free.app]()

