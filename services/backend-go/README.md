# Hello World Api build with Golang

<img src="https://cdn.worldvectorlogo.com/logos/golang-1.svg" width="100" height="100"  text-align="center">


## How to run this service:

run this command, and make sure you are already installed a makefile on your local laptops

to build go binary:
```bash
make -C backend-go
```
to run unit testing:
```bash
make -C backend-go test
```

to run applications:
```bash
make -C backend-go run
```
and it will redirected to your browser `http://localhost:8080`


## OR


You want to build as docker container to run in your local / kubernetes cluster:

make sure you are in `./backend-go` directory / folder:

Build image:
```bash
docker build -t backend-go:1.0 .
```

example output of success build:


```log
❯ docker build -t backend-go:1.0 .
[+] Building 11.0s (9/9) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                                                     0.0s
 => => transferring dockerfile: 36B                                                                                                                                                                      0.0s
 => [internal] load .dockerignore                                                                                                                                                                        0.0s
 => => transferring context: 2B                                                                                                                                                                          0.0s
 => [internal] load metadata for docker.io/library/golang:latest                                                                                                                                         1.2s
 => [internal] load build context                                                                                                                                                                        0.0s
 => => transferring context: 120B                                                                                                                                                                        0.0s
 => [1/4] FROM docker.io/library/golang:latest@sha256:7b297d9abee021bab9046e492506b3c2da8a3722cbf301653186545ecc1e00bb                                                                                   0.0s
 => CACHED [2/4] WORKDIR /app                                                                                                                                                                            0.0s
 => [3/4] COPY . .                                                                                                                                                                                       0.0s
 => [4/4] RUN go build -o main .                                                                                                                                                                         9.2s
 => exporting to image                                                                                                                                                                                   0.4s
 => => exporting layers                                                                                                                                                                                  0.4s
 => => writing image sha256:c9e4beac6c74f453faf1a7026874cd23530e7c44eed4f3c5f418d1fb4e6d5868                                                                                                             0.0s
 => => naming to docker.io/library/backend-go:1.0
```


After build is success you can run this command to start:


```bash
❯ docker run backend-go:1.0 -p 8080:8080
```


