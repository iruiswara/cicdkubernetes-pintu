# CI / CD Agents Explanations

Directory Structure:

```
├── Deployment
│   ├── Jenkinsfile
│   └── README.md
└── ci-agent
    └── README.md
    └── agent.yaml
```

`Deployment` folder contains Jenkinsfile and this Readme docs

`ci-agent  ` folder contains an agent to run ci/cd

---

`agent.yaml`

1. Build Agents:
```yml
...
    - name: kaniko
      image: gcr.io/kaniko-project/executor:debug
      imagePullPolicy: IfNotPresent
      resources: {}
      command:
      - /busybox/cat
      tty: true
      volumeMounts:
        - name: jenkins-docker-cfg
          mountPath: /kaniko/.docker
...
```
i use kaniko as build agent, because It is designed to be used in environments where running privileged Docker daemons or mounting the Docker socket is not desired or allowed. Kaniko doesn't require any special permissions, making it suitable for building container images in Kubernetes clusters.
and i was add `volumeMounts` to mount `kubernetes secret` its give `deployment agent` ability to push docker images to private `image registry`

2. Unit Test Agents:
```yaml
...
    - name: unit-test
      image: xdevbase/go-nodejs:v1.14-v12
      command:
      - cat
      tty: true
      imagePullPolicy: IfNotPresent
      resources: {}
...
```
i use `xdevbase/go-nodejs` for unit testing container because its already contain golang `version 1.14` and `Node.js version 12` available


3. Deployment Agent:
```yml
...
    - name: helm
      image: alpine/helm:3.1.2
      command:
      - cat
      tty: true
      imagePullPolicy: IfNotPresent
      resources: {}
...
```
Next is helm-agent is a container that has Helm installed. 
The Helm charts for backend-go and backend-node are installed using the `helm upgrade --install` command. We can adjust the chart paths, namespace, and other parameters based on our actual Helm charts and deployment configurations.