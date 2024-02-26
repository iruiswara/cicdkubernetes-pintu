# Hello World Api build with nodeJs

<img src="https://cdn.pixabay.com/photo/2015/04/23/17/41/node-js-736399_1280.png" width="100" height="100"  text-align="center">


## How to run this service:


Build the code with command:

```bash

npm install

```

After build run this command:

```bash
node app.js
```


and it will redirected to your browser `http://localhost:3000`


## OR


You want to build as docker container to run in your local / kubernetes cluster:

make sure you are in `./backend-node` directory / folder:

Build image:
```bash
docker build -t backend-node:1.0 .
```

example output of success build:


```log
❯ docker build -t backend-node:1.0 .
[+] Building 8.0s (9/9) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                                                     0.0s
 => => transferring dockerfile: 36B                                                                                                                                                                      0.0s
 => [internal] load .dockerignore                                                                                                                                                                        0.0s
 => => transferring context: 2B                                                                                                                                                                          0.0s
 => [internal] load metadata for docker.io/library/node:14-alpine                                                                                                                                        1.1s
 => [1/4] FROM docker.io/library/node:14-alpine@sha256:434215b487a329c9e867202ff89e704d3a75e554822e07f3e0c0f9e606121b33                                                                                  0.0s
 => [internal] load build context                                                                                                                                                                        0.2s
 => => transferring context: 2.21MB                                                                                                                                                                      0.2s
 => CACHED [2/4] WORKDIR /app                                                                                                                                                                            0.0s
 => [3/4] COPY . .                                                                                                                                                                                       0.1s
 => [4/4] RUN npm install                                                                                                                                                                                6.3s
 => exporting to image                                                                                                                                                                                   0.2s
 => => exporting layers                                                                                                                                                                                  0.1s
 => => writing image sha256:ee5c690012a2639e0f89bfabef7ff356bbde1c13fee9df358402c09ee94db013                                                                                                             0.0s
 => => naming to docker.io/library/backend-node:1.0
```


After build is success you can run this command to start:


```bash
❯ docker run backend-go:1.0 -p 8080:8080
```


