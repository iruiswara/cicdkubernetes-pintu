# CI / CD Explanations

Directory Structure:

```
├── Deployment
│   ├── Jenkinsfile
│   └── README.md
```

`Deployment` folder contains Jenkinsfile and this Readme docs

---

## IMPORTANT SETTINGS:

for Kubernetes Agent settings:

+ [CI-AGENT](../ci-agent/README.md)
  
Create jenkins job:
1. add new item on jenkins
2. choose `Pipelines` for type
3. check on `This project is parameterized`
4. add `String parameter` and add these parameter
    + `repoUrl` > [https://${URL to private repository}]()
    + `registryUrl` >[https://${URL to private registry}]()
    + `BRANCH_NAME` > the branch to deploy it can be (dev,staging,sandbox, production) branch
    + `SERVICE_LANG` as `Choice Parameter` > this parameter to define for which language you want to deploy `go` OR `nodeJs`
    + `NS` > define namespace in kubernetes 
    + `deployDir` > folder to store deployment manifest and settings
---
## Pipeline explaination

default parameters:
```
pipeline {
    environment{
        VERSION="0.0.${BUILD_NUMBER}"
        msName = ""     
    }
...
```
in this section i define default environment for `VERSION` and `msName`, where `VERSION` is used for image `tag` which `${BUILD_NUMBER}` is a jenkins environment variables to get build-number.

stage `Pre Deployment`
```
...
        stage('Pre Deployment'){
            steps{
                script{
                    if("${SERVICE_LANG}" == "go" ){
                        return msName = "backend-go"
                    }else if("${SERVICE_LANG}" == "node"){
                        return msName = "backend-node"
                    }else{
                        return msName = "backend-go,backend-node"
                    }
                }
            }
        }
...
```
this stage is to define variables `msName` because that variables default is `null` 


stage `UNIT TESTING:`

```
        stage('Unit Testing'){
            steps{
                container(name: 'unit-test'){
                    git credentialsId: 'devops-cicd-github', url: '${repoUrl}', branch: "${BRANCH_NAME}"
                   script{
                       if ("${SERVICE_LANG}" == "go" || "${SERVICE_LANG}" == "ALL"){
                           sh "make -C backend-go test"
                       }else if("${SERVICE_LANG}" == "go" || "${SERVICE_LANG}" == "ALL"){
                           sh "make -C backend-node test"
                       }
                   }
                }
            }
        }
```
in this stage we are doing `unit-test` used for testing technique where individual units or components of a software application are tested in isolation from the rest of the system.

stage `Build & push images`

```
stage('Build & push Images'){

            parallel{
                stage('Build go service'){
                    agent { kubernetes { 
                        cloud 'kubernetes' 
                        yamlFile 'node/build.yaml' 
                        } 
                    }
                    steps{
                        container(name: 'kaniko', shell: '/busybox/sh'){
                        git credentialsId: 'devops-cicd-github', url: '${repoUrl}', branch: "${BRANCH_NAME}"
                            script{
                                if("${SERVICE_LANG}" == "go" || "${SERVICE_LANG}" == "ALL"){
                                    sh ''' #!/busybox/sh -xe
                                    /kaniko/executor --context `pwd` --dockerfile ./services/backend-go/Dockerfile  --destination ${registryUrl}/production/${msName}:${VERSIONS} --destination ${registryUrl}/production/${msName}:latest --cleanup && mkdir -p /workspace 
                                    '''
                                }else {
                                    echo "${SERVICE_LANG} not match in condition"
                                    exit 1;
                                }
                            }
                        }
                    }
                }

                stage('Build nodejs service'){
                    steps{
                        container(name: 'kaniko', shell: '/busybox/sh'){
                        git credentialsId: 'devops-cicd-github', url: '${repoUrl}', branch: "${BRANCH_NAME}"
                            script{
                               if("${SERVICE_LANG}" == "node" || "${SERVICE_LANG}" == "ALL") {
                                    sh ''' #!/busybox/sh -xe
                                    /kaniko/executor --context `pwd` --dockerfile ./services/backend-node/Dockerfile --destination ${registryUrl}/production/${msName}:${VERSIONS} --destination ${registryUrl}/production/${msName}:latest --cleanup && mkdir -p /workspace 
                                    '''
                                }else{
                                    echo "${SERVICE_LANG} not match in condition"
                                    exit 1;
                                }
                            }
                        }
                    }
                }
                
            }
        }
```

in this stage i define some settings that allow developer to choose on which language they want to deploy it can `go` services or `node` services and they can choose `both` services to build container images

stage `Deployment`

```
stage('Deployment'){
            environment{
                MS_NAME = "${msName}"
            }
            steps{
               container('helm'){
                script{
                    if("${SERVIICE_LANG}" == "ALL"){
                        def serviceList = SERVICE_NAMES.split(',')
                        def msNameGo = serviceList[0].trim()
                        def msNameNode = serviceList[1].trim()

                    echo "Deploying Service Name (Go): ${msNameGo}"
                    sh"""
                       helm upgrade ${msNameGo} ${deployDir} \
                       --set-string image.repository=${registryUrl}/production/${msName},image.tag=${VERSION} \
                       -f ./${deployDir}/values.yaml --debug --install --namespace ${NS}
                    """

                    echo "Service Name (Node): ${msNameNode}"
                    sh"""
                       helm upgrade ${msNameNode} ${deployDir} \
                       --set-string image.repository=${registryUrl}/production/${msName},image.tag=${VERSION} \
                       -f ./${deployDir}/values.yaml --debug --install --namespace ${NS}
                    """

                    }else{
                    sh"""
                       helm upgrade ${msName} ${deployDir} \
                       --set-string image.repository=${registryUrl}/production/${msName},image.tag=${VERSION} \
                       -f ./${deployDir}/values.yaml --debug --install --namespace ${NS}
                    """
                      }
                    }
                }
            }
```
this stage is to run deployment based on `${SERVICE_LANG}` or `both services`
```
...
            environment{
                MS_NAME = "${msName}"
            }
...
```
this section is added because we run different container so we need to re-define `variables` again

```
...
                    if("${SERVIICE_LANG}" == "ALL"){
                        def serviceList = SERVICE_NAMES.split(',')
                        def msNameGo = serviceList[0].trim()
                        def msNameNode = serviceList[1].trim()

                    echo "Deploying Service Name (Go): ${msNameGo}"
...
```
I added this logic code to separate the variables that we have defined in the `pre deployment stage`, because the `msName` variable contains (,) comma separated so I need to separate the contents of that variable and we can continue deploying with sequential operations.

section `notification`

```
...
            post{
                failure{
                    
                    slackSend channel: '#deployment-channel', color: '#db833a', message: "Failed to start deployment: ${JOB_NAME}", teamDomain: 'company-channel', tokenCredentialId: 'jenkins-deployment-bot'
                }
                success{
                    
                    slackSend channel: '#deployment-channel', color: '#db833a', message: "Deployment Finish for ${JOB_NAME} \nEnvironment: ${deployEnv}", teamDomain: 'company-channel', tokenCredentialId: 'jenkins-deployment-bot'
                }
            }
        }
...
```
this section used for notify developer / user to tell them that their deployment is success / fails